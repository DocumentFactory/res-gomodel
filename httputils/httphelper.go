package httputils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pnocera/res-gomodel/config"
	"github.com/pnocera/res-gomodel/types"
)

//Config struct using viper
type HttpHelper struct {
	conf *config.Config
}

//New Create a new config
func NewHttpHelper(conf *config.Config) *HttpHelper {
	c := HttpHelper{
		conf: conf,
	}

	return &c
}

func (ul *HttpHelper) GetHttpClient() *http.Client {
	return Default()
}

func (ul *HttpHelper) GetDaprUrl(service string, method string) string {
	return fmt.Sprintf(`%s/v1.0/invoke/%s/method/%s`, ul.conf.DaprHostPort(), service, method)
}

func (ul *HttpHelper) DaprGet(service string, method string) ([]byte, error) {

	url := ul.GetDaprUrl(service, method)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// if wfp.User != nil {
	// 	req.Header.Set("X-Forwarded-User", wfp.User.Email)
	// 	req.Header.Set("X-Forwarded-User-Id", wfp.User.ID)
	// 	req.Header.Set("X-Forwarded-User-Ctx", wfp.User.CtxID)
	// }

	client := ul.GetHttpClient()

	return ul.ManageResponse(client.Do(req))

}

//{"errorCode":"ERR_DIRECT_INVOKE","message":"fail to invoke, id: cadsvc, err: failed to invoke target cadsvc after 3 retries"}

func (ul *HttpHelper) ManageResponse(resp *http.Response, err error) ([]byte, error) {

	if err != nil {
		if resp != nil {
			io.Copy(ioutil.Discard, resp.Body)

		}
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var daprerror *types.DaprError
		err2 := json.Unmarshal(body, &daprerror)
		if err2 == nil {
			if daprerror.ErrorCode != "" {
				return nil, errors.New(daprerror.Message)
			}
		}

		return body, nil
	}
	return nil, errors.New("NIL_RESPONSE")

}

func (ul *HttpHelper) DaprPost(service string, method string, data []byte) ([]byte, error) {

	url := ul.GetDaprUrl(service, method)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := ul.GetHttpClient()

	return ul.ManageResponse(client.Do(req))

}
