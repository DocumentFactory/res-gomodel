package types

type EnqueuePayload struct {
	Data   WorkflowData `json:"data"`
	Docs   PListItem    `json:"docs"` // The item to upload / download
	Params interface{}  `json:"params"`
}
