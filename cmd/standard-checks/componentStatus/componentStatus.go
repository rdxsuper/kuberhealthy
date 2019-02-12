// Package componentStatus implements a componentstatus CheckAgent.
package main

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/duration"

	"github.com/Comcast/kuberhealthy/pkg/model"

	// required for oidc kubectl testing
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func init() {
}

func main() {

}

// CheckAgent validates componentstatus objects within the cluster.
type CheckAgent struct {
	model.Check
	FailureTimeStamp map[string]time.Time
	MaxTimeInFailure float64 // TODO - make configurable
	client           *kubernetes.Clientset
}

// New returns a new CheckAgent
func New(client kubernetes.Clientset) model.KuberhealthyCheckServer {
	return &CheckAgent{
		FailureTimeStamp: make(map[string]time.Time),
		MaxTimeInFailure: 300,
		client:           &client,
		Check: model.Check{
			Interval: &duration.Duration{
				Seconds: 120,
			},
			Name: "ComponentStatusCheckAgent",
			Timeout: &duration.Duration{
				Seconds: 60,
			},
			CurrentStatus: &model.Check_Status{
				Errors:  []string{},
				Healthy: true,
			},
		},
	}
}

// Shutdown is implemented to satisfy the KuberhealthyCheck interface, but
// no action is necessary.
func (csc *CheckAgent) Shutdown(context.Context, *model.Request) (*model.Response, error) {
	return nil, nil
}

// Status returns the status of the check as of right now
func (csc *CheckAgent) Status(ctx context.Context, req *model.Request) (*model.Response, error) {
	return &model.Response{
		Check: &csc.Check,
	}, nil
}

// clearErrors clears all errors
func (csc *CheckAgent) clearErrors() {
	csc.Check.CurrentStatus.Errors = []string{}
}

// Start implements the entrypoint for check execution
func (csc *CheckAgent) Start(ctx context.Context, req *model.Request) (*model.Response, error) {
	doneChan := make(chan error)
	// run the check in a goroutine and notify the doneChan when completed
	go func(doneChan chan error) {
		err := csc.doChecks()
		doneChan <- err
	}(doneChan)

	// wait for either a timeout or job completion
	select {
	case <-time.After(durationToTime(*csc.Check.Interval)):
		// The check has timed out because its time to run again
		return nil, errors.New("Failed to complete checks for " + csc.Check.Name + " in time!  Next run came up but check was still running.")
	case <-time.After(durationToTime(*csc.Check.Timeout)):
		// The check has timed out after its specified timeout period
		return nil, errors.New("Failed to complete checks for " + csc.Check.Name + " in time!  Timeout was reached.")
	case err := <-doneChan:
		return nil, err
	}
}

// doChecks executes checks.
func (csc *CheckAgent) doChecks() error {

	// list componentstatuses
	componentList, err := csc.client.CoreV1().ComponentStatuses().List(metav1.ListOptions{})
	if err != nil {
		csc.Check.CurrentStatus.Errors = []string{"Error creating client when checking componentstatuses: " + err.Error()}
		return nil
	}

	// check componentstatuses
	var errorMessages []string
	for _, component := range componentList.Items {
		for _, condition := range component.Conditions {
			currentComponent := component.Name
			// remove the condition from the map if its not still in a failure state
			if condition.Status == v1.ConditionTrue {
				delete(csc.FailureTimeStamp, currentComponent)
				continue
			}
			// if unhealthy...
			timestamp, exists := csc.FailureTimeStamp[currentComponent]
			// add newly failed components to the FailureTimeStamp map
			if !exists {
				csc.FailureTimeStamp[currentComponent] = time.Now()
				continue
			}
			// if a container has been failed for X time, alert
			if time.Now().Sub(timestamp).Seconds() > csc.MaxTimeInFailure {
				errorMessages = append(errorMessages, "componentstatus "+component.Name+" is in a bad state: "+string(condition.Status))
			}

			// loop through all conditions in the failed condition map and
			// remove any that do not exist anymore
			for previouslyFailedComponent := range csc.FailureTimeStamp {
				for _, component := range componentList.Items {
					currentComponent := component.Name
					if currentComponent == previouslyFailedComponent {
						break
					}
				}
				delete(csc.FailureTimeStamp, previouslyFailedComponent)
			}
		}
	}

	// if errors found, set them, if not, clear all
	if len(errorMessages) > 0 {
		csc.Check.CurrentStatus.Errors = errorMessages
	} else {
		csc.clearErrors()
	}

	// nil indicates no system error occurred when checking
	return nil
}

func durationToTime(d duration.Duration) time.Duration {
	return time.Duration(d.GetSeconds()) * time.Second
}
