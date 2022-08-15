package domain

type Login struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"createdAt,omitempty"`
	Token     string `json:"token"`
}
