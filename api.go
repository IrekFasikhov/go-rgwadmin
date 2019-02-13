package rgwadmin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

//API struct for New Client
type API struct {
	endpoint  string
	accessKey string
	secretKey string
	client    *http.Client
}

//NewRGW returns client for Ceph RGW
func NewRGW(endpoint, accessKey, secretKey string) (*API, error) {
	if len(endpoint) < 1 || len(accessKey) < 1 || len(secretKey) < 1 {
		return nil, fmt.Errorf("env RGW_ENDPOINT,RGW_ACCESS_KEY,RGW_SECRET_KEY must be no nil")
	}
	return &API{endpoint, accessKey, secretKey, &http.Client{}}, nil
}

//Query - make request
func (api *API) Query(method, param string, args url.Values) (body []byte, err error) {

	url, err := url.Parse(fmt.Sprintf("%s/admin%s?%s", api.endpoint, param, args.Encode()))
	// fmt.Println("URL:", method, url)
	if err != nil {
		return nil, err
	}

	cred := credentials.NewStaticCredentials(api.accessKey, api.secretKey, "")

	signer := v4.NewSigner(cred)
	signer.DisableRequestBodyOverwrite = true

	re, _ := http.NewRequest(method, url.String(), nil)
	signer.Sign(re, nil, "s3", "default", time.Now().Add(time.Minute))

	resp, err := api.client.Do(re)

	if err != nil {
		return nil, err
	}
	body, err = ioutil.ReadAll(resp.Body)

	return body, err
}

func getReflect(i interface{}, values *url.Values) {

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	for b := 0; b < v.NumField(); b++ {
		v2 := v.Field(b)

		name := t.Field(b).Tag.Get("url")

		for _, name := range strings.Split(name, ",") {

			if v2.Kind() == reflect.Struct {
				getReflect(v2.Interface(), values)
			}

			if v2.Kind() == reflect.Slice {
				for i := 0; i < v2.Len(); i++ {
					item := v2.Index(i)
					getReflect(item.Interface(), values)
				}
			}

			if v2.Kind() == reflect.String ||
				v2.Kind() == reflect.Bool ||
				v2.Kind() == reflect.Int {

				_v2 := fmt.Sprint(v2)
				if len(_v2) > 0 && len(name) > 0 {
					values.Add(name, _v2)
				}
			}

			if v2.Kind() == reflect.Ptr && v2.IsValid() && !v2.IsNil() {
				_v2 := fmt.Sprint(v2.Elem())

				values.Add(name, _v2)
			}
		}
	}

}
