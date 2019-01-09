// Package componentStatus implements a componentstatus checker.
package componentStatus

import (
	"reflect"
	"testing"
	"time"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Checker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Name(t *testing.T) {
	tests := []struct {
		name string
		csc  *Checker
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csc.Name(); got != tt.want {
				t.Errorf("Checker.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_CheckNamespace(t *testing.T) {
	tests := []struct {
		name string
		csc  *Checker
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csc.CheckNamespace(); got != tt.want {
				t.Errorf("Checker.CheckNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Interval(t *testing.T) {
	tests := []struct {
		name string
		csc  *Checker
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csc.Interval(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Checker.Interval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Timeout(t *testing.T) {
	tests := []struct {
		name string
		csc  *Checker
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.csc.Timeout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Checker.Timeout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_Shutdown(t *testing.T) {
	tests := []struct {
		name    string
		csc     *Checker
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.csc.Shutdown(); (err != nil) != tt.wantErr {
				t.Errorf("Checker.Shutdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecker_CurrentStatus(t *testing.T) {
	tests := []struct {
		name  string
		csc   *Checker
		want  bool
		want1 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.csc.CurrentStatus()
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
		csc  *Checker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.csc.clearErrors()
		})
	}
}

func TestChecker_Run(t *testing.T) {
	type args struct {
		client *kubernetes.Clientset
	}
	tests := []struct {
		name    string
		csc     *Checker
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.csc.Run(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("Checker.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChecker_doChecks(t *testing.T) {
	tests := []struct {
		name    string
		csc     *Checker
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.csc.doChecks(); (err != nil) != tt.wantErr {
				t.Errorf("Checker.doChecks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
