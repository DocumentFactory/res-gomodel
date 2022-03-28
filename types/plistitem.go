package types

type PListItem struct {
	Order            int          `json:"Order"`
	ID               string       `json:"ID"`
	DownloadID       string       `json:"DownloadID"`
	UploadLocationID string       `json:"UploadLocationID"`
	ShortName        string       `json:"ShortName"`
	Type             string       `json:"Type"`
	Mimetype         string       `json:"Mimetype"`
	Metadata         string       `json:"Metadata"`
	ParentID         string       `json:"ParentID"`
	Children         []*PListItem `json:"Children"`
	PostMerge        bool         `json:"PostMerge"`
	FileName         string       `json:"FileName"`
	TmpFolderID      string       `json:"tmpFolderID"`
	WFID             string       `json:"wfid"`
	Channel          string       `json:"channel"`
}
