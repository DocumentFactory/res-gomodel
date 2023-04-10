package types

// Nodes Nodes
type Nodes struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Extdata     TypeExtData `json:"extdata"`
	Deleted     bool        `json:"deleted"`
	ParentID    string      `json:"parent_id"`
	Order       int         `json:"norder"`
	Nodetype    string      `json:"nodetype_id"`
	Depth       int         `json:"depth"`
	RootUrl     string      `json:"rooturl"`
	SiteUrl     string      `json:"siteurl"`
	RelativeUrl string      `json:"relativeurl"`
	Ancestors   []*Nodes    `json:"ancestors"`
	Children    []*Nodes    `json:"children"`
}

// TypeExtData Extdata structure
type TypeExtData struct {
	Sub  map[string]interface{} `json:"sub"`
	Meta map[string]interface{} `json:"meta"`
	Tags []string               `json:"tags"`
}

// SPConnectorMeta Sharepoint Connector Metadata
type SPConnectorMeta struct {
	Clientid     string `json:"client_id"`
	Clientsecret string `json:"client_secret"`
}

// SPConnectorMeta Sharepoint Connector Metadata
type S3ConnectorMeta struct {
	EndPoint     string `json:"endpoint"`
	Region       string `json:"region"`
	Clientid     string `json:"client_id"`
	Clientsecret string `json:"client_secret"`
}

// DbConnectorMeta Dropbox Connector Metadata
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
	BucketID        string `json:"bucketid"`
	OutputID        string `json:"outputid"`
	DocumentName    string `json:"documentname"`
	SameFolder      bool   `json:"samefolder"`
}

type ServiceNameData struct {
	Service   string `json:"service"`
	Component string `json:"component"`
}

func FindAllByNodetype(root *Nodes, objtype string) []*Nodes {
	result := make([]*Nodes, 0)
	queue := make([]*Nodes, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.Nodetype == objtype {
			result = append(result, nextUp)
		}
		if len(nextUp.Children) > 0 {
			for _, child := range nextUp.Children {
				results2 := FindAllByNodetype(child, objtype)
				if len(results2) > 0 {
					result = append(result, results2...)
				}
			}
		}
	}
	return result
}

// HasChildrenType HasChildrenType
func (n *Nodes) HasChildrenType(typeName string) bool {
	return len(FindAllByNodetype(n, typeName)) > 0
}
