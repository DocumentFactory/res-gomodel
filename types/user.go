package types

type User struct {
	ID      string      `json:"id"`
	CtxID   string      `json:"ctxid"`
	ExtData UserExtdata `json:"extdata"`
}

type UserExtdata struct {
	Email      string `json:"email"`
	Identities []struct {
		ID       string `json:"id"`
		Provider string `json:"provider"`
		Email    string `json:"email,omitempty"`
	} `json:"identities"`
}
