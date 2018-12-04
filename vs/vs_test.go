package vs

import (
	"testing"
	"errors"
)

func TestSetVaultAddr(t *testing.T) {
	vsConfig := Initialize()

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
			t.Errorf("Unexpected error response for SetVaultAddress(%s) test case. Got error, but expected none", testcase.addr)
		}
		if ((reterr != nil && testcase.err == nil) || (reterr == nil && testcase.err != nil)) {
			t.Errorf("Unexpected error response %v for SetVaultAddress test case with address %s: Expected error %v",
				reterr, testcase.addr, testcase.err)
		}
		if reterr == nil && testcase.err != nil {
			t.Errorf("Unexpected error response for SetVaultAddress(%s) test case. Expected an error, but did notr get one", testcase.addr)
		}
	}
}
