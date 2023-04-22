package types

type WFKeyVal struct {
	RUNID          string   `json:"runid"`
	InstanceType   string   `json:"instancetype"`
	DocThreadSizes []uint64 `json:"docthreadsizes"`
	Status         int      `json:"status"`
}
