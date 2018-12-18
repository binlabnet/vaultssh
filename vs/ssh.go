package vs

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	BUFSIZE = 1024
)

func (vsConfig *VSConfig) openForContent(filename, permissions string, len int) (f *os.File, err error) {
	f, err = os.Create(vsConfig.GetLocalPath())
	return f, err

}

func (vsConfig *VSConfig) saveContent(w *os.File, content []byte) (err error) {
	_, err = w.Write(content)
	return err
}

func (vsConfig *VSConfig) copyFrom(session *ssh.Session) (err error) {
	errchan := make(chan error)
	err = vsConfig.copyFromAux(session, errchan)
	if err != nil {
		return err
	}
	err = <-errchan
	return err
}

func (vsConfig *VSConfig) copyFromAux(session *ssh.Session, errchan chan error) (err error) {

	go func() {
		w, err1 := session.StdinPipe()
		if err1 != nil {
			msg := fmt.Sprintf("Unable to connect to session stdin pipe; %v\n", err1)
			log.Printf(msg)
			errchan <- err1
			return
		}
		r, err2 := session.StdoutPipe()
		if err2 != nil {
			msg := fmt.Sprintf("Unable to connect to session stdout pipe; %v\n", err2)
			log.Printf(msg)
			errchan <- err2
			return
		}

		if err1 == nil && err2 == nil {
			fmt.Fprint(w, "\x00") // start the conversation - request the header

			buf := make([]byte, 1)
			n, err := io.ReadFull(r, buf)
			if err != nil {
				msg := fmt.Sprintf("Unable to read; %v\n", err)
				log.Printf(msg)
				errchan <- err
				return
			}

			code := buf[0]
			switch code {
			case 'C':
				// Receive file code

				buf = make([]byte, BUFSIZE)
				n, err = r.Read(buf)
				if err != nil {
					msg := fmt.Sprintf("Unable to read; %v\n", err)
					log.Printf(msg)
					errchan <- err
					return
				} else {
					s := string(buf[:n]) // Ex: s = "0644 15 HELLO"

					parts := strings.Split(s, " ")

					permissions := parts[0]
					lenstr := parts[1]
					filename := parts[2]

					contentLength, err := strconv.Atoi(lenstr)
					if err != nil {
						msg := fmt.Sprintf("Unable to convert header content length for \"%s\"; %v\n", s, err)
						log.Printf(msg)
						errchan <- err
						return
					}

					if contentLength < BUFSIZE {
						buf = make([]byte, contentLength)
					}

					fmt.Fprint(w, "\x00") // Request for content

					f, err := vsConfig.openForContent(filename, permissions, contentLength)
					if err != nil {
						msg := fmt.Sprintf("Unable to open file %s for content; %v\n", s, err)
						log.Printf(msg)
						errchan <- err
						return
					}
					defer f.Close()

					remainder := contentLength
					for {
						if remainder == 0 {
							break // Assert it was previouisly saved
						} else if remainder < BUFSIZE {
							buf = make([]byte, remainder)
						} else if remainder >= BUFSIZE {
							buf = make([]byte, BUFSIZE)
						}

						n, err := r.Read(buf)
						if err != nil {
							if err == io.EOF {
								err = vsConfig.saveContent(f, buf[:n])
								if err != nil {
									msg := fmt.Sprintf("Unable to save content for \"%s\"; %v\n", s, err)
									log.Printf(msg)
									errchan <- err
									return
								}
								break
							}
							msg := fmt.Sprintf("Unable to read content for \"%s\"; %v\n", s, err)
							log.Printf(msg)
							errchan <- err
							return
						}
						remainder -= n

						err = vsConfig.saveContent(f, buf[:n])
						if err != nil {
							msg := fmt.Sprintf("Unable to save content for \"%s\"; %v\n", s, err)
							log.Printf(msg)
							errchan <- err
							return
						}
					}
				}
			case 'T':
				msg := fmt.Sprintf("Got code T\n")
				log.Printf(msg)
			case 'D':
				msg := fmt.Sprintf("Got code D\n")
				log.Printf(msg)
			case 'E':
				msg := fmt.Sprintf("Got code E\n")
				log.Printf(msg)
			default:
				msg := fmt.Sprintf("Unknown code %v\n", code)
				log.Printf(msg)
			}

			fmt.Fprint(w, "\x00")
		}
		errchan <- nil
	}()

	remote := vsConfig.GetRemotePath()
	cmd := "/usr/bin/scp -qf " + remote
	err = session.Run(cmd)
	if err != nil {
		msg := fmt.Sprintf("Remote scp command \"%s\" failed;  %+v", cmd, err)
		log.Printf(msg)
	}
	return err
}

func (vsConfig *VSConfig) copyTo(session *ssh.Session) (err error) {
	file, err := os.Open(vsConfig.GetLocalPath())
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	errchan := make(chan error)
	err = vsConfig.copyToAux(session, file, stat.Size(), errchan)
	if err != nil {
		return err
	}
	err = <-errchan
	return err
}

func (vsConfig *VSConfig) copyToReader(session *ssh.Session, fileReader io.Reader) (err error) {
	contents_bytes, _ := ioutil.ReadAll(fileReader)
	bytes_reader := bytes.NewReader(contents_bytes)

	errchan := make(chan error)
	err = vsConfig.copyToAux(session, bytes_reader, int64(len(contents_bytes)), errchan)
	if err != nil {
		return err
	}
	err = <-errchan
	return err
}

func getLocalPermString(filename string) (permstr string, err error) {
	perm := os.FileMode(0777)
	flag := os.O_RDONLY
	f, err := os.OpenFile(filename, flag, perm)
	if err != nil {
		return permstr, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return permstr, err
	}
	perm &= fi.Mode()
	permstr = fmt.Sprintf("0%o", perm)
	return permstr, err
}

func (vsConfig *VSConfig) copyToAux(session *ssh.Session, r io.Reader, size int64, errchan chan error) (err error) {
	local := vsConfig.GetLocalPath()
	permissions, err := getLocalPermString(local)
	if err != nil {
		msg := fmt.Sprintf("Unable to get local file permisions for %s; %v\n", local, err)
		log.Printf(msg)
		return err
	}

	filename := path.Base(local)

	remote := vsConfig.GetRemotePath()
	if strings.HasSuffix(remote, "/") {
		remote = remote + filename
	}

	filename = path.Base(remote)

	go func() {
		w, err := session.StdinPipe()
		if err != nil {
			msg := fmt.Sprintf("Unable to connect to session stdin pipe; %v\n", err)
			log.Printf(msg)
			errchan <- err
			return
		} else {
			fmt.Fprintln(w, "C"+permissions, size, filename)
			io.Copy(w, r)
			fmt.Fprint(w, "\x00")
			w.Close()
		}
		errchan <- nil
	}()

	directory := path.Dir(remote)
	cmd := "/usr/bin/scp -qt " + directory
	err = session.Run(cmd)
	if err != nil {
		msg := fmt.Sprintf("Remote scp command \"%s\" failed;  %+v", cmd, err)
		log.Printf(msg)
		return err
	}

	return err
}

func (vsConfig *VSConfig) setupSession() (session *ssh.Session, err error) {
	err = vsConfig.VaultLogin()
	if err != nil {
		return session, err
	}

	clientConfig, err := vsConfig.getSignedCertConfig()
	if err != nil {
		return session, err
	}

	addr := fmt.Sprintf("%s:%d", vsConfig.GetSshServerHost(), vsConfig.GetSshServerPort())

	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		msg := fmt.Sprintf("Unable to dial %s: %v\n", addr, err)
		log.Printf(msg)
		return session, err
	}

	session, err = client.NewSession()
	if err != nil {
		msg := fmt.Sprintf("Unable to create session: %v\n", err)
		log.Printf(msg)
		return session, err
	}
	return session, err
}

func (vsConfig *VSConfig) ScpSessionAux() (err error) {

	session, err := vsConfig.setupSession()
	if err != nil {
		return err
	}
	defer session.Close()

	if vsConfig.GetMode() == SCPTO {
		err = vsConfig.copyTo(session)
	} else if vsConfig.GetMode() == SCPFROM {
		err = vsConfig.copyFrom(session)
		if err != nil {
			return err
		}
	} // TODO: adjust SetMode validation
	return err
}

func (vsConfig *VSConfig) StartSessionAux() (err error) {
	session, err := vsConfig.setupSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	terminalModes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	fileDescriptor := int(os.Stdin.Fd())

	if terminal.IsTerminal(fileDescriptor) {

		originalState, err := terminal.MakeRaw(fileDescriptor)
		if err != nil {
			msg := fmt.Sprintf("request for raw terminal failed: %v\n", err)
			log.Printf(msg)
			return err
		}
		defer terminal.Restore(fileDescriptor, originalState)

		termWidth, termHeight, err := terminal.GetSize(fileDescriptor)
		if err != nil {
			msg := fmt.Sprintf("request for terminal size failed: %v\n", err)
			log.Printf(msg)
			return err
		}

		err = session.RequestPty(vsConfig.GetTermType(), termHeight, termWidth, terminalModes)
		if err != nil {
			msg := fmt.Sprintf("request for pseudo terminal failed: %v\n", err)
			log.Printf(msg)
			return err
		}
	}

	if err := session.Shell(); err != nil {
		msg := fmt.Sprintf("failed to start shell: %v\n", err)
		log.Printf(msg)
		return err
	}

	session.Wait() // This call blocks until the user exits the session (e.g. via CTRL + D)

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
