package types

type SPItem struct {
	CreatedDateTime      string                 `json:"createdDateTime"`
	LastModifiedDateTime string                 `json:"lastModifiedDateTime"`
	Id                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	DisplayName          string                 `json:"displayName"`
	WebUrl               string                 `json:"webUrl"`
	Folder               map[string]interface{} `json:"folder"`
	File                 map[string]interface{} `json:"file"`
}
