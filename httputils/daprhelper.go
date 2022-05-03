package httputils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/pnocera/res-gomodel/config"
	"github.com/pnocera/res-gomodel/enums"
	"github.com/pnocera/res-gomodel/logs"
	"github.com/pnocera/res-gomodel/types"
	"go.uber.org/zap"
)

type DaprHelper struct {
	conf   *config.Config
	client dapr.Client
	logh   *logs.LogHelper
}

func NewDaprHelper(conf *config.Config) *DaprHelper {
	c := DaprHelper{
		conf: conf,
	}
	client, err := dapr.NewClientWithAddress(conf.DaprHostPort())
	if err != nil {
		panic(err)
	}

	c.client = client

	c.logh = logs.NewLogHelper(conf)

	return &c
}

func (ul *DaprHelper) Get(ctx context.Context, servicename string, api string) ([]byte, error) {
	response, err := ul.client.InvokeMethod(ctx, servicename, api, "get")
	if err != nil {
		ul.logh.Error("Dapr Error getting request ", zap.String("servicename", servicename), zap.String("api", api), zap.String("error", err.Error()))
	}
	return response, err
}

func (ul *DaprHelper) Post(ctx context.Context, servicename string, api string, data interface{}) ([]byte, error) {
	response, err := ul.client.InvokeMethodWithCustomContent(ctx, servicename, api, "post", "application/json", data)
	if err != nil {
		ul.logh.Error("Dapr Error posting request ", zap.String("servicename", servicename), zap.String("api", api), zap.String("error", err.Error()))
	}
	return response, err
}

func (ul *DaprHelper) GetExtension(ctx context.Context, Mimetype string) (string, error) {

	response, err := ul.client.InvokeMethod(ctx, enums.DataSvc, "mimes/extension/"+strings.ReplaceAll(Mimetype, "/", "|"), "get")

	if err != nil {
		ul.logh.Error("Dapr Error getting extension ", zap.String("mimetype", Mimetype), zap.String("error", err.Error()))
		return "", err
	}

	return string(response), nil

}

func (ul *DaprHelper) GetNodeSecret(ctx context.Context, node types.Nodes) (map[string]interface{}, error) {
	return ul.GetSecret(ctx, node.Nodetype, node.ID)
}

func (ul *DaprHelper) GetSecret(ctx context.Context, nodetype string, secretid string) (map[string]interface{}, error) {
	secretbytes, err := ul.Post(ctx, enums.SecretSvc, "secretv1", types.SecretData{
		Action:     "read",
		SecretPath: fmt.Sprintf("%s/%s", ul.conf.DFEnv(), nodetype),
		SecretID:   secretid,
	})

	if err != nil {
		ul.logh.Error("Dapr Error getting secret ",
			zap.String("nodetype", nodetype),
			zap.String("secretid", secretid),
			zap.String("error", err.Error()))
		return nil, err
	}

	var secretresp types.SecretData

	err = json.Unmarshal(secretbytes, &secretresp)

	if err != nil {
		ul.logh.Error("Dapr Error unmarshalling secret ",
			zap.String("nodetype", nodetype),
			zap.String("secretid", secretid),
			zap.String("bytes", string(secretbytes)),
			zap.String("error", err.Error()))
		return nil, err
	}

	if !secretresp.Ok {
		err = errors.New(secretresp.ErrMsg)
		ul.logh.Error("Dapr Error getting secret ",
			zap.String("nodetype", nodetype),
			zap.String("secretid", secretid),
			zap.String("error", err.Error()))

		return nil, err
	}

	return secretresp.SecretValue, nil
}

func (ul *DaprHelper) GetService(ctx context.Context, filename string) (string, error) {

	ext := filepath.Ext(filename)
	if ext == "" {
		ext = "docx"
	}

	response, err := ul.client.InvokeMethod(ctx, enums.DataSvc, "mimes/serviceext/"+ext, "get")
	if err != nil {
		ul.logh.Error("Dapr Error getting service ",
			zap.String("service", enums.DataSvc),
			zap.String("extension", ext),
			zap.String("error", err.Error()))

		return "", err
	}

	var svcmeta types.ServiceNameData

	err = json.Unmarshal(response, &svcmeta)
	if err != nil {
		ul.logh.Error("Dapr Error unmarshalling service ",
			zap.String("service", enums.DataSvc),
			zap.String("extension", ext),
			zap.String("response", string(response)),
			zap.String("error", err.Error()))
		return "", err
	}

	return svcmeta.Service, nil

}
