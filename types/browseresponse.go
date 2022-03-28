package types

type BrowseResponse struct {
	CurrentPage int      `json:"currentPage"` //the current page
	PerPage     int      `json:"perPage"`     //number of results per page
	LastPage    int      `json:"lastPage"`    //last page
	Data        []*Nodes `json:"data"`        //the nodes result
	Total       int      `json:"total"`       //the total number of nodes
}
