package types

type WFKeyVal struct {
	RunID        string `json:"runid"`
	WorkflowType int    `json:"workflowtype"`
	TotalSize    uint64 `json:"totalsize"`
	HasMerge     bool   `json:"hasmerge"`
	NumDocs      int    `json:"numdocs"`
	Status       int    `json:"status"`
}
