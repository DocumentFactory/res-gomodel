package types

type DaprError struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}
