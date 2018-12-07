package vs

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
)

func (vsConfig *VSConfig) StartSessionAux() (err error) {
	err = vsConfig.VaultLogin()
	if err != nil {
		return err
	}

	clientConfig, err := vsConfig.getSignedCertConfig()
	addr := fmt.Sprintf("%s:%d", vsConfig.GetSshServerHost(), vsConfig.GetSshServerPort())

	conn, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		msg := fmt.Sprintf("Unable to connect to %s: %v\n", addr, err)
		log.Printf(msg)
		return err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		msg := fmt.Sprintf("Unable to create session: %v\n", err)
		log.Printf(msg)
		return err
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	in, _ := session.StdinPipe()

	terminalModes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty(vsConfig.GetTermType(), vsConfig.GetTermRows(), vsConfig.GetTermCols(), terminalModes); err != nil {
		msg := fmt.Sprintf("request for pseudo terminal failed: %v\n", err)
		log.Printf(msg)
		return err
	}

	if err := session.Shell(); err != nil {
		msg := fmt.Sprintf("failed to start shell: %v\n", err)
		log.Printf(msg)
		return err
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
		if strings.HasPrefix(str, "logout") {
			break
		}
	}
	return err // No error; Ready to exit
}

func (vsConfig *VSConfig) getSignedCertConfig() (clientConfig *ssh.ClientConfig, err error) {

	pubkey, privkey, err := vsConfig.VaultReadSSHKey()
	if err != nil {
		return clientConfig, err
	}

	signedCrt, err := vsConfig.SignPubKey(pubkey)
	authorizedKeysBytes := []byte(signedCrt)
	privkeyBytes := []byte(privkey)

	pcert, _, _, _, err := ssh.ParseAuthorizedKey(authorizedKeysBytes)
	if err != nil {
		return clientConfig, err
	}

	upkey, err := ssh.ParseRawPrivateKey(privkeyBytes)
	if err != nil {
		return clientConfig, err
	}

	usigner, err := ssh.NewSignerFromKey(upkey)
	if err != nil {
		return clientConfig, err
	}

	ucertSigner, err := ssh.NewCertSigner(pcert.(*ssh.Certificate), usigner)

	if err != nil {
		return clientConfig, err
	}

	clientConfig = &ssh.ClientConfig{
		User:            vsConfig.GetSshUsername(),
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(ucertSigner)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return clientConfig, err
}
