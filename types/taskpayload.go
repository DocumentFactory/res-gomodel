package types

type TaskPayload struct {
	Data   WorkflowData `json:"data"`
	Docs   PListItem    `json:"docs"` // The item to upload / download
	Params interface{}  `json:"params"`
}

type TaskResponse struct {
	TASKID string `json:"taskid"`
}
