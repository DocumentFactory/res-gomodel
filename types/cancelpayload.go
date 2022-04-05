package types

type CancelPayload struct {
	WFID  string `json:"wfid"`
	RUNID string `json:"runid"`
}
