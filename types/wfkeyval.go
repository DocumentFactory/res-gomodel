package types

type WFKeyVal struct {
	RUNID          string   `json:"runid"`
	DocThreadSizes []uint64 `json:"docthreadsizes"`
	Status         int      `json:"status"`
}
