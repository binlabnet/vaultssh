package vs

import (
	"reflect"
	"testing"

	"github.com/hashicorp/vault/api"
)

var vsc = Initialize().(*VSConfig)

func TestVSConfig_GetSigningRole(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetSigningRole(); got != tt.want {
				t.Errorf("VSConfig.GetSigningRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetSigningRole(t *testing.T) {
	type args struct {
		role string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetSigningRole(tt.args.role)
		})
	}
}

func TestVSConfig_GetMode(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetMode(); got != tt.want {
				t.Errorf("VSConfig.GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetMode(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetMode(tt.args.mode)
		})
	}
}

func TestVSConfig_GetVaultAddress(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetVaultAddress(); got != tt.want {
				t.Errorf("VSConfig.GetVaultAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetVaultAddress(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := vsc.SetVaultAddress(tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.SetVaultAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVSConfig_GetPublicKeyPath(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetPublicKeyPath(); got != tt.want {
				t.Errorf("VSConfig.GetPublicKeyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetPublicKeyPath(t *testing.T) {
	type args struct {
		keypath string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetPublicKeyPath(tt.args.keypath)
		})
	}
}

func TestVSConfig_GetPrivateKeyPath(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetPrivateKeyPath(); got != tt.want {
				t.Errorf("VSConfig.GetPrivateKeyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetPrivateKeyPath(t *testing.T) {
	type args struct {
		keypath string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetPrivateKeyPath(tt.args.keypath)
		})
	}
}

func TestVSConfig_GetSshServerHost(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetSshServerHost(); got != tt.want {
				t.Errorf("VSConfig.GetSshServerHost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetSshServerHost(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetSshServerHost(tt.args.host)
		})
	}
}

func TestVSConfig_GetSshServerPort(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetSshServerPort(); got != tt.want {
				t.Errorf("VSConfig.GetSshServerPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetSshServerPort(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetSshServerPort(tt.args.port)
		})
	}
}

func TestVSConfig_GetTermType(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetTermType(); got != tt.want {
				t.Errorf("VSConfig.GetTermType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetTermType(t *testing.T) {
	type args struct {
		termtype string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetTermType(tt.args.termtype)
		})
	}
}

func TestVSConfig_GetSshUsername(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetSshUsername(); got != tt.want {
				t.Errorf("VSConfig.GetSshUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetSshUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetSshUsername(tt.args.username)
		})
	}
}

func TestVSConfig_GetUsername(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetUsername(); got != tt.want {
				t.Errorf("VSConfig.GetUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetUsername(tt.args.username)
		})
	}
}

func TestVSConfig_GetPasswd(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetPasswd(); got != tt.want {
				t.Errorf("VSConfig.GetPasswd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetPasswd(t *testing.T) {
	type args struct {
		pw string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetPasswd(tt.args.pw)
		})
	}
}

func TestVSConfig_GetVaultClient(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     *api.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetVaultClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VSConfig.GetVaultClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetVaultClient(t *testing.T) {
	type args struct {
		client *api.Client
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetVaultClient(tt.args.client)
		})
	}
}

func TestVSConfig_GetVaultToken(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetVaultToken(); got != tt.want {
				t.Errorf("VSConfig.GetVaultToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetVaultToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetVaultToken(tt.args.token)
		})
	}
}

func TestVSConfig_GetPrivateKey(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetPrivateKey(); got != tt.want {
				t.Errorf("VSConfig.GetPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetPrivateKey(t *testing.T) {
	type args struct {
		privKey string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetPrivateKey(tt.args.privKey)
		})
	}
}

func TestVSConfig_GetPublicKey(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		want     string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vsConfig.GetPublicKey(); got != tt.want {
				t.Errorf("VSConfig.GetPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVSConfig_SetPublicKey(t *testing.T) {
	type args struct {
		pubKey string
	}
	tests := []struct {
		name     string
		vsConfig *VSConfig
		args     args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vsConfig.SetPublicKey(tt.args.pubKey)
		})
	}
}

func TestVSConfig_AddKeyPair(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vsConfig.AddKeyPair(); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.AddKeyPair() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVSConfig_SignPubKey(t *testing.T) {
	type args struct {
		pubKey string
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
			gotSignedCrt, err := tt.vsConfig.SignPubKey(tt.args.pubKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.SignPubKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSignedCrt != tt.wantSignedCrt {
				t.Errorf("VSConfig.SignPubKey() = %v, want %v", gotSignedCrt, tt.wantSignedCrt)
			}
		})
	}
}

func TestVSConfig_StartSession(t *testing.T) {
	tests := []struct {
		name     string
		vsConfig *VSConfig
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.vsConfig.StartSession(); (err != nil) != tt.wantErr {
				t.Errorf("VSConfig.StartSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
