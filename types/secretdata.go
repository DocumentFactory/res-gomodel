package types

type SecretData struct {
	Action      string                 `json:"action"`
	SecretPath  string                 `json:"secretpath"`
	SecretID    string                 `json:"secretid"`
	SecretValue map[string]interface{} `json:"secretvalue"`
	Ok          bool                   `json:"ok"`
	ErrMsg      string                 `json:"errmsg"`
}
