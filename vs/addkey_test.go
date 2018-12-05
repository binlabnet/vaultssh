package vs

import "testing"

func TestVSConfig_AddKeyPairAux(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vsConfig.AddKeyPairAux(); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.AddKeyPairAux() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
