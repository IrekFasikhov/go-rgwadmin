package rgwadmin

import (
	"encoding/json"
)

//Quota struct
type Quota struct {
	UID        string  `json:"user_id" url:"uid"`
	QuotaType  string  `url:"quota-type"`
	Enabled    *bool   `json:"enabled" url:"quota,enabled"`
	CheckOnRaw bool    `json:"check_on_raw"`
	MaxSize    *uint64 `json:"max_size" url:"max-size"`
	MaxSizeKb  *uint64 `json:"max_size_kb" url:"max-size-kb"`
	MaxObjects *uint64 `json:"max_objects" url:"max-objects"`
}

//GetQuota for bucket and user
func (api *API) GetQuota(quota Quota, param string) (*Quota, error) {
	quota.QuotaType = param
	body, err := api.Query("GET", "/user", GetValues(quota))
	if err != nil {
		return nil, err
	}

	ref := &Quota{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//SetQuota for bucket and user
func (api *API) SetQuota(quota Quota, param string) (*Quota, error) {
	quota.QuotaType = param
	body, err := api.Query("PUT", "/user", GetValues(quota))
	if err != nil {
		return nil, err
	}

	ref := &Quota{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}
