package storage

type VaultStorage struct{}

func (v VaultStorage) put(key string, value string) error {
	panic("implement me")
}

func (v VaultStorage) get(key string) (string, error) {
	panic("implement me")
}
