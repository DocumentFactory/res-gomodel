package types

type WFKeyVal struct {
	RUNID        string `json:"runid"`
	InstanceType string `json:"instancetype"`
	Status       int    `json:"status"`
}
