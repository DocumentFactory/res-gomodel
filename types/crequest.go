package types

type CRequest struct {
	Docs      *PListItem    `json:"docs"`
	Wfrequest *WorkflowData `json:"wfrequest"`
}
