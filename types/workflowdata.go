package types

type WorkflowData struct {
	Builder         *Nodes     `json:"buildernode"`
	Book            *Nodes     `json:"booknode"`
	InputConnector  *Nodes     `json:"inputconnector"`
	OutputConnector *Nodes     `json:"outputconnector"`
	PostData        *PListItem `json:"postdata"`
	User            *User      `json:"user"`
	BaseUrl         string     `json:"baseurl"`
	WFID            string     `json:"wfid"`
	RUNID           string     `json:"runid"`
	InstanceType    int        `json:"instancetype"`
	TotalSize       uint64     `json:"totalsize"`
	NumDocs         int        `json:"numdocs"`
	HasMerge        bool       `json:"hasmerge"`
}

type WorkflowInfo struct {
	Builder         *Nodes `json:"buildernode"`
	Book            *Nodes `json:"booknode"`
	InputConnector  *Nodes `json:"inputconnector"`
	OutputConnector *Nodes `json:"outputconnector"`
	User            *User  `json:"user"`
	BaseUrl         string `json:"baseurl"`
	WFID            string `json:"wfid"`
	RUNID           string `json:"runid"`
	InstanceType    int    `json:"instancetype"`
	TotalSize       uint64 `json:"totalsize"`
	NumDocs         int    `json:"numdocs"`
	HasMerge        bool   `json:"hasmerge"`
}
