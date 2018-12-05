package vs

import (
	"reflect"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestVSConfig_StartSessionAux(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vsConfig.StartSessionAux(); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.StartSessionAux() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVSConfig_getSignedCertConfig(t *testing.T) {
	tests := []struct {
		name             string
		vsConfig         *VSConfig
		wantClientConfig *ssh.ClientConfig
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClientConfig, err := tt.vsConfig.getSignedCertConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.getSignedCertConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClientConfig, tt.wantClientConfig) {
				t.Errorf("VSConfig.getSignedCertConfig() = %v, want %v", gotClientConfig, tt.wantClientConfig)
			}
		})
	}
}
