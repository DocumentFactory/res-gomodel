package types

type ProcessedRequest struct {
	WFID   string       `json:"wfid"`
	RUNID  string       `json:"runid"`
	Signal string       `json:"signal"`
	Docs   []*PListItem `json:"docs"`
	Ok     bool         `json:"ok"`
	ErrMsg string       `json:"errmsg"`
}
