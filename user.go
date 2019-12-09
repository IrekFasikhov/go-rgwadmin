package rgwadmin

import (
	"encoding/json"
	"net/url"
)

//User struct
type User struct {
	Tenant      string        `json:"tenant,omitempty" url:"tenant"`
	UID         string        `json:"user_id" url:"uid"`
	DisplayName string        `json:"display_name" url:"display-name"`
	Email       string        `json:"email" url:"email"`
	Suspended   *int          `json:"suspended" url:"suspended"`
	MaxBuckets  *int          `json:"max_buckets" url:"max-buckets"`
	Subusers    []interface{} `json:"subusers"`
	Keys        []Key         `json:"keys"`
	SwiftKeys   []interface{} `json:"swift_keys"`
	Caps        []Cap         `json:"caps"`
	PurgeData   *int          `url:"purge-data"`
	GenerateKey *bool         `url:"generate-key"`
	// Quota       Quota
}

//Key struct
type Key struct {
	User      string `json:"user"`
	AccessKey string `json:"access_key" url:"access-key"`
	SecretKey string `json:"secret_key" url:"secret-key"`
}

//Cap struct
type Cap struct {
	Type string `json:"type"`
	Perm string `json:"perm"`
}

//GetUsers - list all RGW users
func (api *API) GetUsers() (*[]string, error) {
	body, err := api.Query("GET", "/metadata/user", nil)
	if err != nil {
		return nil, err
	}
	var ref *[]string
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//GetUser - http://docs.ceph.com/docs/mimic/radosgw/adminops/#get-user-info
func (api *API) GetUser(user User) (*User, error) {
	body, err := api.Query("GET", "/user", GetValues(user))
	if err != nil {
		return nil, err
	}
	ref := &User{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//CreateUser - http://docs.ceph.com/docs/mimic/radosgw/adminops/#create-user
func (api *API) CreateUser(user User) (*User, error) {
	body, err := api.Query("PUT", "/user", GetValues(user))
	if err != nil {
		return nil, err
	}

	ref := &User{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//RemoveUser - http://docs.ceph.com/docs/mimic/radosgw/adminops/#remove-user
func (api *API) RemoveUser(user User) error {
	_, err := api.Query("DELETE", "/user", GetValues(user))
	return err
}

//ModifyUser - http://docs.ceph.com/docs/mimic/radosgw/adminops/#modify-user
func (api *API) ModifyUser(user User) (*User, error) {

	body, err := api.Query("POST", "/user", GetValues(user))
	if err != nil {
		return nil, err
	}

	ref := &User{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//GetValues reflect
func GetValues(i interface{}) url.Values {

	values := url.Values{}
	getReflect(i, &values)
	return values
}
