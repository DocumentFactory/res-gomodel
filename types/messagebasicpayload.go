package types

type BasicPayload struct {
	ID   string       `json:"ID"`
	Data WorkflowData `json:"data"`
}
