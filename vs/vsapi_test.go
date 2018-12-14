package vs

import (
	"reflect"
	"testing"

	"github.com/hashicorp/vault/api"
)

func TestAddKeyPair(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if err := AddKeyPair(tt.args.vsapi); (err != nil) != tt.wantErr {
				t.Errorf("AddKeyPair() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStartSession(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if err := StartSession(tt.args.vsapi); (err != nil) != tt.wantErr {
				t.Errorf("StartSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignPubKey(t *testing.T) {
	type args struct {
		vsapi  VsApi
		pubKey string
	}
	tests := []struct {
		name          string
		args          args
		wantSignedCrt string
		wantErr       bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignedCrt, err := SignPubKey(tt.args.vsapi, tt.args.pubKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignPubKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSignedCrt != tt.wantSignedCrt {
				t.Errorf("SignPubKey() = %v, want %v", gotSignedCrt, tt.wantSignedCrt)
			}
		})
	}
}

func TestVaultReadSSHKey(t *testing.T) {
	type args struct {
		vsapi VsApi
	}
	tests := []struct {
		name        string
		args        args
		wantPubKey  string
		wantPrivKey string
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPubKey, gotPrivKey, err := VaultReadSSHKey(tt.args.vsapi)
			if (err != nil) != tt.wantErr {
				t.Errorf("VaultReadSSHKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPubKey != tt.wantPubKey {
				t.Errorf("VaultReadSSHKey() gotPubKey = %v, want %v", gotPubKey, tt.wantPubKey)
			}
			if gotPrivKey != tt.wantPrivKey {
				t.Errorf("VaultReadSSHKey() gotPrivKey = %v, want %v", gotPrivKey, tt.wantPrivKey)
			}
		})
	}
}

func TestVaultWriteSSHKey(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if err := VaultWriteSSHKey(tt.args.vsapi); (err != nil) != tt.wantErr {
				t.Errorf("VaultWriteSSHKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVaultLogin(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if err := VaultLogin(tt.args.vsapi); (err != nil) != tt.wantErr {
				t.Errorf("VaultLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetSigningRole(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetSigningRole(tt.args.vsapi); got != tt.want {
				t.Errorf("GetSigningRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSigningRole(t *testing.T) {
	type args struct {
		vsapi VsApi
		role  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSigningRole(tt.args.vsapi, tt.args.role)
		})
	}
}

func TestGetMode(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetMode(tt.args.vsapi); got != tt.want {
				t.Errorf("GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMode(t *testing.T) {
	type args struct {
		vsapi VsApi
		mode  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMode(tt.args.vsapi, tt.args.mode)
		})
	}
}

func TestGetVaultAddress(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetVaultAddress(tt.args.vsapi); got != tt.want {
				t.Errorf("GetVaultAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetVaultAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"SetVaultAddress-1", args{"http://hostabc.acme.com:8200"}, false},
		{"SetVaultAddress-2", args{"https://hostabc.acme.com:8200"}, false},
		{"SetVaultAddress-3", args{"http://hostabc.acme.com"}, true},
		{"SetVaultAddress-4", args{"httpfoo://hostabc.acme.com:8200"}, true},
		{"SetVaultAddress-5", args{"http://:8200"}, true},
	}
	vsc := NewVSConfig()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetVaultAddress(vsc, tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.SetVaultAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPublicKeyPath(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetPublicKeyPath(tt.args.vsapi); got != tt.want {
				t.Errorf("GetPublicKeyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPublicKeyPath(t *testing.T) {
	type args struct {
		vsapi   VsApi
		keypath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPublicKeyPath(tt.args.vsapi, tt.args.keypath)
		})
	}
}

func TestGetPrivateKeyPath(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetPrivateKeyPath(tt.args.vsapi); got != tt.want {
				t.Errorf("GetPrivateKeyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPrivateKeyPath(t *testing.T) {
	type args struct {
		vsapi   VsApi
		keypath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPrivateKeyPath(tt.args.vsapi, tt.args.keypath)
		})
	}
}

func TestGetSshServerHost(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetSshServerHost(tt.args.vsapi); got != tt.want {
				t.Errorf("GetSshServerHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSshServerHost(t *testing.T) {
	type args struct {
		vsapi VsApi
		host  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSshServerHost(tt.args.vsapi, tt.args.host)
		})
	}
}

func TestGetSshServerPort(t *testing.T) {
	type args struct {
		vsapi VsApi
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSshServerPort(tt.args.vsapi); got != tt.want {
				t.Errorf("GetSshServerPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSshServerPort(t *testing.T) {
	type args struct {
		vsapi VsApi
		port  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSshServerPort(tt.args.vsapi, tt.args.port)
		})
	}
}

func TestGetTermType(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetTermType(tt.args.vsapi); got != tt.want {
				t.Errorf("GetTermType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetTermType(t *testing.T) {
	type args struct {
		vsapi    VsApi
		termtype string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetTermType(tt.args.vsapi, tt.args.termtype)
		})
	}
}

func TestGetSshUsername(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetSshUsername(tt.args.vsapi); got != tt.want {
				t.Errorf("GetSshUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSshUsername(t *testing.T) {
	type args struct {
		vsapi    VsApi
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSshUsername(tt.args.vsapi, tt.args.username)
		})
	}
}

func TestGetUsername(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetUsername(tt.args.vsapi); got != tt.want {
				t.Errorf("GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetUsername(t *testing.T) {
	type args struct {
		vsapi    VsApi
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetUsername(tt.args.vsapi, tt.args.username)
		})
	}
}

func TestGetPasswd(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetPasswd(tt.args.vsapi); got != tt.want {
				t.Errorf("GetPasswd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPasswd(t *testing.T) {
	type args struct {
		vsapi VsApi
		pw    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPasswd(tt.args.vsapi, tt.args.pw)
		})
	}
}

func TestGetVaultClient(t *testing.T) {
	type args struct {
		vsapi VsApi
	}
	tests := []struct {
		name string
		args args
		want *api.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVaultClient(tt.args.vsapi); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVaultClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetVaultClient(t *testing.T) {
	type args struct {
		vsapi  VsApi
		client *api.Client
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetVaultClient(tt.args.vsapi, tt.args.client)
		})
	}
}

func TestGetVaultToken(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetVaultToken(tt.args.vsapi); got != tt.want {
				t.Errorf("GetVaultToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetVaultToken(t *testing.T) {
	type args struct {
		vsapi VsApi
		token string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetVaultToken(tt.args.vsapi, tt.args.token)
		})
	}
}

func TestGetPrivateKey(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetPrivateKey(tt.args.vsapi); got != tt.want {
				t.Errorf("GetPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPrivateKey(t *testing.T) {
	type args struct {
		vsapi   VsApi
		privKey string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPrivateKey(tt.args.vsapi, tt.args.privKey)
		})
	}
}

func TestGetPublicKey(t *testing.T) {
	type args struct {
		vsapi VsApi
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
			if got := GetPublicKey(tt.args.vsapi); got != tt.want {
				t.Errorf("GetPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPublicKey(t *testing.T) {
	type args struct {
		vsapi  VsApi
		pubKey string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetPublicKey(tt.args.vsapi, tt.args.pubKey)
		})
	}
}
