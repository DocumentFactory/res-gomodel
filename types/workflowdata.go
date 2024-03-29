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
	InstanceType    string `json:"instancetype"`
}
