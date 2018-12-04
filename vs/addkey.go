package vs

func (vsConfig *VSConfig) AddKeyPairAux() (err error) {
	err = vsConfig.VaultLogin()
	if err != nil {
		return err
	}
	err = vsConfig.VaultWriteSSHKey()
	if err != nil {
		return err
	}
	return err
}
