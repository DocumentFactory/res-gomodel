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

//HasChildrenType HasChildrenType
func (n *PListItem) HasChildrenType(typeName string) bool {
	return len(FindAllByItemtype(n, typeName)) > 0
}

//FindAllByItemtype FindAllByItemtype
func FindAllByItemtype(root *PListItem, objtype string) []*PListItem {
	result := make([]*PListItem, 0)
	queue := make([]*PListItem, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.Type == objtype {
			result = append(result, nextUp)
		}
		if len(nextUp.Children) > 0 {
			for _, child := range nextUp.Children {
				results2 := FindAllByItemtype(child, objtype)
				if len(results2) > 0 {
					result = append(result, results2...)
				}

			}
		}
	}
	return result
}
