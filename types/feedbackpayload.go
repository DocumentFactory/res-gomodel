package types

type FeedbackPayload struct {
	WFID   string       `json:"wfid"`
	RUNID  string       `json:"runid"`
	TASKID string       `json:"taskid"`
	Docs   []*PListItem `json:"docs"`
	Ok     bool         `json:"ok"`
	ErrMsg string       `json:"errmsg"`
}
