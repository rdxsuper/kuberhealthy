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
package khstatecrd

import (
	"reflect"
	"testing"

	"github.com/Comcast/kuberhealthy/pkg/health"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestKuberhealthyState_String(t *testing.T) {
	tests := []struct {
		name string
		h    KuberhealthyState
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.String(); got != tt.want {
				t.Errorf("KuberhealthyState.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKuberhealthyState_DeepCopyInto(t *testing.T) {
	type args struct {
		out *KuberhealthyState
	}
	tests := []struct {
		name string
		h    KuberhealthyState
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.DeepCopyInto(tt.args.out)
		})
	}
}

func TestKuberhealthyState_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name string
		h    KuberhealthyState
		want runtime.Object
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.DeepCopyObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KuberhealthyState.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKuberhealthyState(t *testing.T) {
	type args struct {
		name string
		spec health.CheckDetails
	}
	tests := []struct {
		name string
		args args
		want KuberhealthyState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKuberhealthyState(tt.args.name, tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKuberhealthyState() = %v, want %v", got, tt.want)
			}
		})
	}
}
