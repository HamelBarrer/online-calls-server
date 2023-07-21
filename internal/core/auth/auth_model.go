package auth

type Auth struct {
	Username string `json:"username,omitempty" required:"true"`
	Password string `json:"password,omitempty" required:"true"`
}
