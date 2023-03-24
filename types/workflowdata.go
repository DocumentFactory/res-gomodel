package types

type WorkflowData struct {
	//Component either docb or publ
	Builder         *Nodes     `json:"buildernode"`
	Book            *Nodes     `json:"booknode"`
	InputConnector  *Nodes     `json:"inputconnector"`
	OutputConnector *Nodes     `json:"outputconnector"`
	PostData        *PListItem `json:"postdata"`
	User            *User      `json:"user"`
	BaseUrl         string     `json:"baseurl"`
	WFID            string     `json:"wfid"`
	RUNID           string     `json:"runid"`
	WorkflowType    int        `json:"workflowtype"`
}
