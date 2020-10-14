package storage

import (
	"errors"
	"github.com/hashicorp/vault/api"
)

type VaultStorageConfig struct {
	ClientConfig *api.Config
	Token        string
	SecretPath   string
}

type VaultStorage struct {
	client     *api.Client
	secretPath string
}

func NewVaultStorage(vaultConfig *VaultStorageConfig) (*VaultStorage, error) {
	apiClient, err := api.NewClient(vaultConfig.ClientConfig)
	if err != nil {
		return nil, err
	}

	apiClient.SetToken(vaultConfig.Token)
	return &VaultStorage{
		client:     apiClient,
		secretPath: vaultConfig.SecretPath,
	}, nil
}

func (v *VaultStorage) ChangeSecretPath(secretPath string) {
	v.secretPath = secretPath
}

func (v *VaultStorage) put(key string, value string) error {
	_, err := v.client.Logical().Write(v.secretPath, map[string]interface{}{
		key: value,
	})

	if err != nil {
		return err
	}

	return nil
}

func (v *VaultStorage) get(key string) (string, error) {
	secret, err := v.client.Logical().Read(v.secretPath)
	if err != nil {
		return "", err
	}

	value, ok := secret.Data[key]
	if !ok {
		return "", errors.New("key could not be found")
	}

	return value.(string), nil
}
