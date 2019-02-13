package rgwadmin

import (
	"encoding/json"
)

//Usage struct
type Usage struct {
	Entries []struct {
		User    string `json:"user"`
		Buckets []struct {
			Bucket     string `json:"bucket"`
			Time       string `json:"time"`
			Epoch      uint64 `json:"epoch"`
			Owner      string `json:"owner"`
			Categories []struct {
				Category      string `json:"category"`
				BytesSent     uint64 `json:"bytes_sent"`
				BytesReceived uint64 `json:"bytes_received"`
				Ops           uint64 `json:"ops"`
				SuccessfulOps uint64 `json:"successful_ops"`
			} `json:"categories"`
		} `json:"buckets"`
	} `json:"entries"`
	Summary []struct {
		User       string `json:"user"`
		Categories []struct {
			Category      string `json:"category"`
			BytesSent     uint64 `json:"bytes_sent"`
			BytesReceived uint64 `json:"bytes_received"`
			Ops           uint64 `json:"ops"`
			SuccessfulOps uint64 `json:"successful_ops"`
		} `json:"categories"`
		Total struct {
			BytesSent     uint64 `json:"bytes_sent"`
			BytesReceived uint64 `json:"bytes_received"`
			Ops           uint64 `json:"ops"`
			SuccessfulOps uint64 `json:"successful_ops"`
		} `json:"total"`
	} `json:"summary"`
	Start       string `url:"start"` //Example:	2012-09-25 16:00:00
	End         string `url:"end"`
	ShowEntries *bool  `url:"show-entries"`
	ShowSummary *bool  `url:"show-summary"`
	RemoveAll   *bool  `url:"remove-all"` //true
}

//GetUsage - http://docs.ceph.com/docs/mimic/radosgw/adminops/#get-usage
func (api *API) GetUsage(usage Usage) (*Usage, error) {
	body, err := api.Query("GET", "/usage", GetValues(usage))
	if err != nil {
		return nil, err
	}
	ref := &Usage{}
	err = json.Unmarshal(body, &ref)
	return ref, err
}

//TrimUsage - http://docs.ceph.com/docs/mimic/radosgw/adminops/#trim-usage
func (api *API) TrimUsage(usage Usage) error {
	_, err := api.Query("DELETE", "/usage", GetValues(usage))
	return err
}
