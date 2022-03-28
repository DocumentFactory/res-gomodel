package types

type Message struct {
	ID      string      `json:"id"`
	Payload interface{} `json:"payload"`
}
