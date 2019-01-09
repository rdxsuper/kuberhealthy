// Package podStatus implements a pod health checker for Kuberhealthy.  Pods are checked
// to ensure they are not restarting too much and are in a healthy lifecycle
// phase.
package podStatus

import (
	"reflect"
	"testing"
	"time"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func TestNew(t *testing.T) {
	type args struct {
		namespace string
	}
	tests := []struct {
		name string
		args args
		want *Checker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Name(t *testing.T) {
	tests := []struct {
		name string
		psc  *Checker
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.psc.Name(); got != tt.want {
				t.Errorf("Checker.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_CheckNamespace(t *testing.T) {
	tests := []struct {
		name string
		psc  *Checker
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.psc.CheckNamespace(); got != tt.want {
				t.Errorf("Checker.CheckNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Interval(t *testing.T) {
	tests := []struct {
		name string
		psc  *Checker
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.psc.Interval(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Checker.Interval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Timeout(t *testing.T) {
	tests := []struct {
		name string
		psc  *Checker
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.psc.Timeout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Checker.Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Shutdown(t *testing.T) {
	tests := []struct {
		name    string
		psc     *Checker
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.psc.Shutdown(); (err != nil) != tt.wantErr {
				t.Errorf("Checker.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecker_CurrentStatus(t *testing.T) {
	tests := []struct {
		name  string
		psc   *Checker
		want  bool
		want1 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.psc.CurrentStatus()
			if got != tt.want {
				t.Errorf("Checker.CurrentStatus() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Checker.CurrentStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestChecker_clearErrors(t *testing.T) {
	tests := []struct {
		name string
		psc  *Checker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.psc.clearErrors()
		})
	}
}

func TestChecker_Run(t *testing.T) {
	type args struct {
		client *kubernetes.Clientset
	}
	tests := []struct {
		name    string
		psc     *Checker
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.psc.Run(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("Checker.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecker_doChecks(t *testing.T) {
	tests := []struct {
		name    string
		psc     *Checker
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.psc.doChecks(); (err != nil) != tt.wantErr {
				t.Errorf("Checker.doChecks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecker_podFailures(t *testing.T) {
	tests := []struct {
		name         string
		psc          *Checker
		wantFailures []string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFailures, err := tt.psc.podFailures()
			if (err != nil) != tt.wantErr {
				t.Errorf("Checker.podFailures() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFailures, tt.wantFailures) {
				t.Errorf("Checker.podFailures() = %v, want %v", gotFailures, tt.wantFailures)
			}
		})
	}
}

func Test_componentFailures(t *testing.T) {
	type args struct {
		client kubernetes.Interface
	}
	tests := []struct {
		name         string
		args         args
		wantFailures []string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFailures, err := componentFailures(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("componentFailures() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFailures, tt.wantFailures) {
				t.Errorf("componentFailures() = %v, want %v", gotFailures, tt.wantFailures)
			}
		})
	}
}
