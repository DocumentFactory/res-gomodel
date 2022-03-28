package types

//Nodes Nodes
type Nodes struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Extdata  TypeExtData `json:"extdata"`
	Deleted  bool        `json:"deleted"`
	ParentID string      `json:"parent_id"`
	Order    int         `json:"norder"`
	Nodetype string      `json:"nodetype_id"`
	Children []*Nodes    `json:"children"`
}

//TypeExtData Extdata structure
type TypeExtData struct {
	Sub  map[string]interface{} `json:"sub"`
	Meta map[string]interface{} `json:"meta"`
	Tags []string               `json:"tags"`
}

//SPConnectorMeta Sharepoint Connector Metadata
type SPConnectorMeta struct {
	Clientid     string `json:"client_id"`
	Clientsecret string `json:"client_secret"`
}

//DbConnectorMeta Dropbox Connector Metadata
type DbConnectorMeta struct {
	Token string `json:"token"`
}

type BactConverterMeta struct {
	Outputformat string `json:"outputformat"`
}

type BookLevelMeta struct {
	TemplateConnector string `json:"templateconnector"`
	TemplateId        string `json:"templateid"`
	Split             bool   `json:"split"`
	Toc               bool   `json:"toc"`
	TocLevels         int    `json:"toclevels"`
	TocDocs           bool   `json:"tocdocs"`
	FileName          string `json:"filename"`
}

type BactSetHeadersMeta struct {
	HeaderTemplateConnector string `json:"headertemplateconnector"`
	HeaderTemplate          string `json:"headertemplate"`
	FileName                string `json:"filename"`
	Mimetype                string
}

type BactSetFooterMeta struct {
	FooterTemplateConnector string `json:"footerstemplateconnector"`
	FooterTemplate          string `json:"footerstemplate"`
	FileName                string `json:"filename"`
	Mimetype                string
}

type HasFilename struct {
	FileName string
}

type BactOutputMeta struct {
	OutputConnector string `json:"outputconnector"`
	OutputID        string `json:"outputid"`
	DocumentName    string `json:"documentname"`
	SameFolder      bool   `json:"samefolder"`
}

type ServiceNameData struct {
	Service   string `json:"service"`
	Component string `json:"component"`
}
