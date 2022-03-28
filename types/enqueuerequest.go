package types

type EnqueueRequest struct {
	Service string      `json:"service"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}
