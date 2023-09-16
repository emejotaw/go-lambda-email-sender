package types

type EmailRequestDTO struct {
	Email    string `json:"email"`
	Type     string `json:"type"`
	Provider string `json:"provider"`
}
