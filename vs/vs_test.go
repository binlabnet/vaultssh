package vs

import (
	"testing"
	"errors"
)

var vsConfig = Initialize()

func TestSetVaultAddr(t *testing.T) {

	err := errors.New("Illegal address")

	tables := []struct {
		addr string
		err error
	}{
		{"http://hostabc.acme.com:8200", nil},
		{"https://hostabc.acme.com:8200", nil},
		{"http://hostabc.acme.com", err},
		{"https://hostabc.acme.com", err},
		{"https://:", err},
	}

	for _, testcase := range(tables) {
		reterr := SetVaultAddress(vsConfig, testcase.addr)
		if reterr != nil && testcase.err == nil {
			t.Errorf("Unexpected error response for SetVaultAddress %s test case. Got an error, but expected none", testcase.addr)
		}
		if reterr == nil && testcase.err != nil {
			t.Errorf("Unexpected error response for SetVaultAddress(%s) test case. Expected an error, but did notr get one", testcase.addr)
		}
	}
}

func TestGetVaultAddr(t *testing.T) {

	tables := []struct {
		addrset string
		addrget string
	}{
		{"http://hostabc.acme.com:8200", "http://hostabc.acme.com:8200"},
		{"https://hostabc.acme.com:8200", "https://hostabc.acme.com:8200"},
	}

	for _, testcase := range(tables) {
		reterr := SetVaultAddress(vsConfig, testcase.addrset)
		if reterr != nil {
			t.Errorf("Unexpected error response from the GetVaultAddress %s test case. Got sn error, but expected none", testcase.addrset)
		}
		retaddr := GetVaultAddress(vsConfig)
		if retaddr != testcase.addrget {
			t.Errorf("Unexpected response for GetVaultAddress %s test case. The returned address %s does not match the expected", testcase.addrset, testcase.addrget)
		}
	}
}
