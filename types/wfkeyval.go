package types

type WFKeyVal struct {
	RunID        string `json:"runid"`
	WorkflowType int    `json:"workflowtype"`
	TotalSize    uint64 `json:"totalsize"`
	Status       int    `json:"status"`
}
