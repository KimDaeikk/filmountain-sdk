package utils

type VaultStorage struct {
	*Storage
}

var vaultStore *VaultStorage

func VaultStore() *VaultStorage {
	return vaultStore
}

func NewVaultStore(filename string) error {
	vaultDefault := map[string]string{
		"id":      "",
		"address": "",
		"tx":      "",
	}

	s, err := NewStorage(filename, vaultDefault, true)
	if err != nil {
		return err
	}

	vaultStore = &VaultStorage{s}

	return nil
}
