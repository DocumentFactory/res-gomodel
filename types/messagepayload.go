package types

type MessagePayload struct {
	ID     string       `json:"ID"`
	Data   WorkflowData `json:"data"`
	Docs   PListItem    `json:"docs"` // The item to upload / download
	Params interface{}  `json:"params"`
}
