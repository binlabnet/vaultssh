package vs

import "testing"

func TestVSConfig_SignPubKeyAux(t *testing.T) {
	type args struct {
		pubkey string
	}
	tests := []struct {
		name          string
		vsConfig      *VSConfig
		args          args
		wantSignedCrt string
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignedCrt, err := tt.vsConfig.SignPubKeyAux(tt.args.pubkey)
			if (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.SignPubKeyAux() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSignedCrt != tt.wantSignedCrt {
				t.Errorf("VSConfig.SignPubKeyAux() = %v, want %v", gotSignedCrt, tt.wantSignedCrt)
			}
		})
	}
}

func TestVSConfig_VaultReadSSHKey(t *testing.T) {
	tests := []struct {
		name        string
		vsConfig    *VSConfig
		wantPubkey  string
		wantPrivkey string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPubkey, gotPrivkey, err := tt.vsConfig.VaultReadSSHKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.VaultReadSSHKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPubkey != tt.wantPubkey {
				t.Errorf("VSConfig.VaultReadSSHKey() gotPubkey = %v, want %v", gotPubkey, tt.wantPubkey)
			}
			if gotPrivkey != tt.wantPrivkey {
				t.Errorf("VSConfig.VaultReadSSHKey() gotPrivkey = %v, want %v", gotPrivkey, tt.wantPrivkey)
			}
		})
	}
}

func TestVSConfig_VaultWriteSSHKey(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vsConfig.VaultWriteSSHKey(); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.VaultWriteSSHKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVSConfig_VaultLogin(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vsConfig.VaultLogin(); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.VaultLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
