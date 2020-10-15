package secret

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

// VaultStore is a Secret Store backed by Hashicorp Vault
type VaultStore struct {
	client     *api.Client
	secretPath string
}

func NewVaultStore(secretPath string) VaultStore {
	return VaultStore{
		secretPath: secretPath,
	}
}

// Connect to the Vault server with a static token (instead of VAULT_TOKEN env variable)
func (s *VaultStore) Connect(url, token string) error {
	client, err := api.NewClient(&api.Config{
		Address: url,
	})
	if err != nil {
		return err
	}

	s.client = client
	s.client.SetToken(token)

	return nil
}

func (s VaultStore) GetTwitchToken() (string, error) {
	secret, err := s.client.Logical().Read(s.secretPath)
	if err != nil {
		return "", fmt.Errorf("ERROR: fetch Twitch Token from Store - %v", err)
	}

	// This can happen from the underlying api
	// @see github.com/hashicorp/vault/api@v1.0.4/logical.go:81
	if secret == nil {
		return "", fmt.Errorf("ERROR: strange things happening - no error and no ptr to secret")
	}

	dataMap := secret.Data["data"].(map[string]interface{})
	token, ok := dataMap["token"]
	if !ok {
		return "", fmt.Errorf("Twitch Token not found in Vault")
	}

	return token.(string), nil
}
