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
package health

import (
	"net/http"
	"reflect"
	"testing"
)

func TestState_AddError(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		h    *State
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.AddError(tt.args.s...)
		})
	}
}

func TestState_WriteHTTPStatusResponse(t *testing.T) {
	type args struct {
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		h       *State
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.WriteHTTPStatusResponse(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("State.WriteHTTPStatusResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewState(t *testing.T) {
	tests := []struct {
		name string
		want State
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewState() = %v, want %v", got, tt.want)
			}
		})
	}
}
