package httputils

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/pnocera/res-gomodel/config"
	"github.com/pnocera/res-gomodel/enums"
	"github.com/pnocera/res-gomodel/types"
)

type DaprHelper struct {
	conf   *config.Config
	client dapr.Client
}

func NewDaprHelper(conf *config.Config) *DaprHelper {
	c := DaprHelper{
		conf: conf,
	}
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	c.client = client

	return &c
}

func (ul *DaprHelper) Get(ctx context.Context, servicename string, api string) ([]byte, error) {
	return ul.client.InvokeMethod(ctx, servicename, api, "get")
}

func (ul *DaprHelper) Post(ctx context.Context, servicename string, api string, data interface{}) ([]byte, error) {
	return ul.client.InvokeMethodWithCustomContent(ctx, servicename, api, "post", "application/json", data)
}

func (ul *DaprHelper) GetExtension(ctx context.Context, Mimetype string) (string, error) {

	response, err := ul.client.InvokeMethod(ctx, enums.DataSvc, "mimes/extension/"+strings.ReplaceAll(Mimetype, "/", "|"), "get")

	if err != nil {
		return "", err
	}

	return string(response), nil

}

func (ul *DaprHelper) GetService(ctx context.Context, filename string) (string, error) {

	ext := filepath.Ext(filename)
	if ext == "" {
		ext = "docx"
	}

	response, err := ul.client.InvokeMethod(ctx, enums.DataSvc, "mimes/serviceext/"+ext, "get")
	if err != nil {
		return "", err
	}

	var svcmeta types.ServiceNameData

	err = json.Unmarshal(response, &svcmeta)
	if err != nil {
		return "", err
	}

	return svcmeta.Service, nil

}
