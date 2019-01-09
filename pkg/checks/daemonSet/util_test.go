package daemonSet

import "testing"

func Test_randString(t *testing.T) {
	type args struct {
		n int
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
			if got := randString(tt.args.n); got != tt.want {
				t.Errorf("randString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHostname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHostname(); got != tt.want {
				t.Errorf("getHostname() = %v, want %v", got, tt.want)
			}
		})
	}
}
