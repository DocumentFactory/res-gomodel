package types

type BrowseRequest struct {
	AuthConfig AuthConfig `json:"authconfig"` // The AuthConfig configuration
	Node       Nodes      `json:"node"`       // The node to browse
	WhereName  string     `json:"where_name"` // The  filter name contains
	Sort       string     `json:"sort"`       // The sort order, eg name_desc
	Page       int        `json:"page"`       // The page number to retrieve
	Limit      int        `json:"limit"`      // The number of items to retrieve

}
