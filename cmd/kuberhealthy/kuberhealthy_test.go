/* Copyright 2018 Comcast Cable Communications Management, LLC
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Comcast/kuberhealthy/pkg/health"
	"k8s.io/client-go/kubernetes"
)

func TestNewKuberhealthy(t *testing.T) {
	type args struct {
		client *kubernetes.Clientset
	}
	tests := []struct {
		name string
		args args
		want *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKuberhealthy(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKuberhealthy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKuberhealthy_setCheckExecutionError(t *testing.T) {
	type args struct {
		checkName string
		err       error
	}
	tests := []struct {
		name string
		k    *Kuberhealthy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.setCheckExecutionError(tt.args.checkName, tt.args.err)
		})
	}
}

func TestKuberhealthy_AddCheck(t *testing.T) {
	type args struct {
		c KuberhealthyCheck
	}
	tests := []struct {
		name string
		k    *Kuberhealthy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.AddCheck(tt.args.c)
		})
	}
}

func TestKuberhealthy_Shutdown(t *testing.T) {
	tests := []struct {
		name string
		k    *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.Shutdown()
		})
	}
}

func TestKuberhealthy_StopChecks(t *testing.T) {
	tests := []struct {
		name string
		k    *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.StopChecks()
		})
	}
}

func TestKuberhealthy_Start(t *testing.T) {
	tests := []struct {
		name string
		k    *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.Start()
		})
	}
}

func TestKuberhealthy_StartChecks(t *testing.T) {
	tests := []struct {
		name string
		k    *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.StartChecks()
		})
	}
}

func TestKuberhealthy_addCheckStopChan(t *testing.T) {
	type args struct {
		checkName string
		stopChan  chan bool
	}
	tests := []struct {
		name string
		k    *Kuberhealthy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.addCheckStopChan(tt.args.checkName, tt.args.stopChan)
		})
	}
}

func TestKuberhealthy_sigChecks(t *testing.T) {
	tests := []struct {
		name string
		k    *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.sigChecks()
		})
	}
}

func TestKuberhealthy_masterStatusMonitor(t *testing.T) {
	type args struct {
		becameMasterChan chan bool
		lostMasterChan   chan bool
	}
	tests := []struct {
		name string
		k    *Kuberhealthy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.masterStatusMonitor(tt.args.becameMasterChan, tt.args.lostMasterChan)
		})
	}
}

func TestKuberhealthy_runCheck(t *testing.T) {
	type args struct {
		stopChan chan bool
		c        KuberhealthyCheck
	}
	tests := []struct {
		name string
		k    *Kuberhealthy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.runCheck(tt.args.stopChan, tt.args.c)
		})
	}
}

func TestKuberhealthy_storeCheckState(t *testing.T) {
	type args struct {
		checkName string
		details   health.CheckDetails
	}
	tests := []struct {
		name    string
		k       *Kuberhealthy
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.storeCheckState(tt.args.checkName, tt.args.details); (err != nil) != tt.wantErr {
				t.Errorf("Kuberhealthy.storeCheckState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKuberhealthy_StartWebServer(t *testing.T) {
	tests := []struct {
		name string
		k    *Kuberhealthy
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.StartWebServer()
		})
	}
}

func TestKuberhealthy_writeHealthCheckError(t *testing.T) {
	type args struct {
		w     http.ResponseWriter
		r     *http.Request
		err   error
		state health.State
	}
	tests := []struct {
		name string
		k    *Kuberhealthy
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.k.writeHealthCheckError(tt.args.w, tt.args.r, tt.args.err, tt.args.state)
		})
	}
}

func TestKuberhealthy_prometheusMetricsHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		k       *Kuberhealthy
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.prometheusMetricsHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Kuberhealthy.prometheusMetricsHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKuberhealthy_healthCheckHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		k       *Kuberhealthy
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.healthCheckHandler(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Kuberhealthy.healthCheckHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKuberhealthy_getCurrentState(t *testing.T) {
	tests := []struct {
		name    string
		k       *Kuberhealthy
		want    health.State
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.k.getCurrentState()
			if (err != nil) != tt.wantErr {
				t.Errorf("Kuberhealthy.getCurrentState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kuberhealthy.getCurrentState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKuberhealthy_getCheck(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		k       *Kuberhealthy
		args    args
		want    KuberhealthyCheck
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.k.getCheck(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Kuberhealthy.getCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kuberhealthy.getCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
