package vault

import (
	"fmt"

	vault "github.com/hashicorp/vault/api"
	"github.com/pnocera/res-gomodel/config"
	"github.com/pnocera/res-gomodel/logs"
	"go.uber.org/zap"
)

type VaultHelper struct {
	conf   *config.Config
	client *vault.Client
	logh   *logs.LogHelper
}

func NewVaultHelper(conf *config.Config) *VaultHelper {
	c := VaultHelper{
		conf: conf,
		logh: logs.NewLogHelper(conf),
	}

	config := vault.DefaultConfig()
	config.Address = conf.VaultHostPort()
	client, err := vault.NewClient(config)
	if err != nil {
		c.logh.Panic(err.Error())

	}

	// client.SetToken(token)
	// client.
	c.client = client

	return &c
}

func (vh *VaultHelper) Getv1(api string) (map[string]interface{}, error) {

	secret, err := vh.client.Logical().Read(api)
	if err != nil {
		vh.logh.Error("Error getting secret ", zap.String("api", api), zap.String("error", err.Error()))
		return nil, err
	}

	if len(secret.Warnings) > 0 {
		for _, sec := range secret.Warnings {
			vh.logh.Warn("Warning getting secret ", zap.String("message", sec))
		}
	}

	vh.logh.Debug(fmt.Sprintf("%v", secret))

	return secret.Data, err
}

func (vh *VaultHelper) Setv1(api string, data map[string]interface{}) error {

	secret, err := vh.client.Logical().Write(api, data)
	if err != nil {
		vh.logh.Error("Error writing secret ", zap.String("api", api), zap.String("error", err.Error()))
		return err
	}

	if len(secret.Warnings) > 0 {
		for _, sec := range secret.Warnings {
			vh.logh.Warn("Warning writing secret ", zap.String("message", sec))
		}
	}

	return err
}

func (vh *VaultHelper) Delete(api string) error {

	secret, err := vh.client.Logical().Delete(api)
	if err != nil {
		vh.logh.Error("Error deleting secret ", zap.String("api", api), zap.String("error", err.Error()))
		return err
	}

	if len(secret.Warnings) > 0 {
		for _, sec := range secret.Warnings {
			vh.logh.Warn("Warning deleting secret ", zap.String("message", sec))
		}
	}

	return err
}
