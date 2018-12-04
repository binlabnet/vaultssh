package main

import (
	"fmt"
	"github.com/richard-mauri/vaultssh/vs"
	"log"
	"os"
)

func main() {

	vsapi := vs.Initialize()

	switch mode := vs.GetMode(vsapi); mode {
	case vs.ADDKEY:
		os.Exit(vs.Addkey(vsapi))
	case vs.SSH:
		os.Exit(vs.Ssh(vsapi))
	default:
		msg := fmt.Sprintf("Illegal mode %s: must be either %s or %s\n", vs.GetMode(vsapi), vs.ADDKEY, vs.SSH)
		log.Printf(msg)
		os.Exit(-1)
	}

	os.Exit(0)
}
