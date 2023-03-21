package enums

// ConnectorSubtype ConnectorSubtype
type ConnectorSubtype string

const (
	SharepointConnector     string = "or.co.ob.sh"
	LocalConnector          string = "or.co.ob.lo"
	SftpConnector           string = "or.co.ob.sf"
	BuilderActionMergeType  string = "ou.ba.mr"
	DropboxConnector        string = "or.co.ob.db"
	BuilderActionOutputType string = "ou.ba.op"
	ConvertNodeType         string = "ou.ba.cv"
	SetHeadersType          string = "ou.ba.hd"
	SetFooterType           string = "ou.ba.ft"
	BookLevelType           string = "ou.bo.lv"
)

type TaskError string

const (
	//TaskDownloadError Task Download Error
	TaskDownloadError = "TaskDownloadError"
	//TaskProcessError Task Process Error
	TaskProcessError = "TaskProcessError"
	//TaskPreProcessError Task PreProcess Error
	TaskPreProcessError = "TaskPreProcessError"
	//TaskMergeError Task Merge Error
	TaskMergeError = "TaskMergeError"
	//TaskUploadError Task Upload Error
	TaskUploadError = "TaskUploadError"
	//TaskUploadError Task Upload Error
	FinalizeError = "FinalizeError"
)

type DaprService string

// services names
const (
	DataSvc       = "datalayer"
	NodeapiSvc    = "apisvc"
	Rbacsvc       = "rbacsvc"
	Compapi       = "compapisvc"
	Qsvc          = "qsvc"
	SecretSvc     = "gosecret"
	SharepointSvc = "spgwsvc"
	SftpSvc       = "sftpgwsvc"
)

// actions names
const (
	ActionPreprocess  = "preprocess"
	ActionProcess     = "process"
	ActionMerge       = "merge"
	ActionBook        = "book"
	ActionPostProcess = "postprocess"
	ActionFinalize    = "finalize"
	ActionDownload    = "download"
	ActionUpload      = "upload"
)
