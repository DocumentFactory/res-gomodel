package types

type AuthConfig struct {
	SiteURL      string `json:"siteurl"`      // SPSite or SPWeb URL, which is the context target for the API calls
	ClientID     string `json:"clientid"`     // Client ID obtained when registering the AddIn
	ClientSecret string `json:"clientsecret"` // Client Secret obtained when registering the AddIn
}
