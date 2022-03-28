package types

type SpUploadParams struct {
	AuthConfig AuthConfig `json:"authconfig"` // The AuthConfig configuration
	Dest       string     `json:"dest"`
	Baseurl    string     `json:"baseurl"`
	Name       string     `json:"name"`
}
