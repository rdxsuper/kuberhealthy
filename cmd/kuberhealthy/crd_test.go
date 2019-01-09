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
	"reflect"
	"testing"

	"github.com/Comcast/kuberhealthy/pkg/health"
	"github.com/Comcast/kuberhealthy/pkg/khstatecrd"
)

func Test_setCheckCRDState(t *testing.T) {
	type args struct {
		checkName string
		client    *khstatecrd.KuberhealthyStateClient
		state     health.CheckDetails
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setCheckCRDState(tt.args.checkName, tt.args.client, tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("setCheckCRDState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sanitizeCRDName(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeCRDName(tt.args.c); got != tt.want {
				t.Errorf("sanitizeCRDName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ensureCRDExists(t *testing.T) {
	type args struct {
		checkName string
		client    *khstatecrd.KuberhealthyStateClient
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ensureCRDExists(tt.args.checkName, tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("ensureCRDExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getCheckCRDState(t *testing.T) {
	type args struct {
		c      KuberhealthyCheck
		client *khstatecrd.KuberhealthyStateClient
	}
	tests := []struct {
		name    string
		args    args
		want    health.CheckDetails
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCheckCRDState(tt.args.c, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCheckCRDState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCheckCRDState() = %v, want %v", got, tt.want)
			}
		})
	}
}
