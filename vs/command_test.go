package vs

import (
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name      string
		wantVsapi VsApi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVsapi := Initialize(); !reflect.DeepEqual(gotVsapi, tt.wantVsapi) {
				t.Errorf("Initialize() = %v, want %v", gotVsapi, tt.wantVsapi)
			}
		})
	}
}

func TestAddkey(t *testing.T) {
	type args struct {
		vsapi VsApi
	}
	tests := []struct {
		name         string
		args         args
		wantExitcode int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExitcode := Addkey(tt.args.vsapi); gotExitcode != tt.wantExitcode {
				t.Errorf("Addkey() = %v, want %v", gotExitcode, tt.wantExitcode)
			}
		})
	}
}

func TestSsh(t *testing.T) {
	type args struct {
		vsapi VsApi
	}
	tests := []struct {
		name         string
		args         args
		wantExitcode int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExitcode := Ssh(tt.args.vsapi); gotExitcode != tt.wantExitcode {
				t.Errorf("Ssh() = %v, want %v", gotExitcode, tt.wantExitcode)
			}
		})
	}
}
