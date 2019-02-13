package rgwadmin

import (
	"encoding/json"
)

//Bucket struct
type Bucket struct {
	Bucket            string `json:"bucket" url:"bucket"`
	Zonegroup         string `json:"zonegroup"`
	PlacementRule     string `json:"placement_rule"`
	ExplicitPlacement struct {
		DataPool      string `json:"data_pool"`
		DataExtraPool string `json:"data_extra_pool"`
		IndexPool     string `json:"index_pool"`
	} `json:"explicit_placement"`
	ID        string `json:"id"`
	Marker    string `json:"marker"`
	IndexType string `json:"index_type"`
	Owner     string `json:"owner"`
	Ver       string `json:"ver"`
	MasterVer string `json:"master_ver"`
	Mtime     string `json:"mtime"`
	MaxMarker string `json:"max_marker"`
	Usage     struct {
		RgwMain struct {
			Size           uint64 `json:"size"`
			SizeActual     uint64 `json:"size_actual"`
			SizeUtilized   uint64 `json:"size_utilized"`
			SizeKb         uint64 `json:"size_kb"`
			SizeKbActual   uint64 `json:"size_kb_actual"`
			SizeKbUtilized uint64 `json:"size_kb_utilized"`
			NumObjects     uint64 `json:"num_objects"`
		} `json:"rgw.main"`
		RgwMultimeta struct {
			Size           uint64 `json:"size"`
			SizeActual     uint64 `json:"size_actual"`
			SizeUtilized   uint64 `json:"size_utilized"`
			SizeKb         uint64 `json:"size_kb"`
			SizeKbActual   uint64 `json:"size_kb_actual"`
			SizeKbUtilized uint64 `json:"size_kb_utilized"`
			NumObjects     uint64 `json:"num_objects"`
		} `json:"rgw.multimeta"`
	} `json:"usage"`
	BucketQuota struct {
		Enabled    *bool   `json:"enabled"`
		CheckOnRaw bool    `json:"check_on_raw"`
		MaxSize    *uint64 `json:"max_size"`
		MaxSizeKb  *uint64 `json:"max_size_kb"`
		MaxObjects *uint64 `json:"max_objects"`
	} `json:"bucket_quota"`
	Police bool `url:"policy"`
}

//Policy struct
type Policy struct {
	ACL struct {
		ACLUserMap []struct {
			User string `json:"user"`
			ACL  int    `json:"acl"`
		} `json:"acl_user_map"`
		ACLGroupMap []interface{} `json:"acl_group_map"`
		GrantMap    []struct {
			ID    string `json:"id"`
			Grant struct {
				Type struct {
					Type int `json:"type"`
				} `json:"type"`
				ID         string `json:"id"`
				Email      string `json:"email"`
				Permission struct {
					Flags int `json:"flags"`
				} `json:"permission"`
				Name    string `json:"name"`
				Group   int    `json:"group"`
				URLSpec string `json:"url_spec"`
			} `json:"grant"`
		} `json:"grant_map"`
	} `json:"acl"`
	Owner struct {
		ID          string `json:"id"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
}

//ListBuckets - http://docs.ceph.com/docs/mimic/radosgw/adminops/#get-bucket-info
func (api *API) ListBuckets() ([]string, error) {
	body, err := api.Query("GET", "/bucket", nil)
	if err != nil {
		return nil, err
	}
	var s []string
	err = json.Unmarshal(body, &s)
	return s, err
}

//InfoBucket - http://docs.ceph.com/docs/mimic/radosgw/adminops/#get-bucket-info
func (api *API) InfoBucket(bucket string) (*Bucket, error) {

	b := Bucket{
		Bucket: bucket,
	}
	body, err := api.Query("GET", "/bucket", GetValues(b))
	if err != nil {
		return nil, err
	}
	ref := &Bucket{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//GetBucketPolicy - http://docs.ceph.com/docs/mimic/radosgw/adminops/#get-bucket-or-object-policy
func (api *API) GetBucketPolicy(bucket string) (*Policy, error) {
	b := Bucket{
		Bucket: bucket,
		Police: true,
	}
	body, err := api.Query("GET", "/bucket", GetValues(b))
	if err != nil {
		return nil, err
	}
	ref := &Policy{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}
