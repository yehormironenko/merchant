package requests

type RegisterUser struct {
	Login    string `json:"login" validate:"required,gte=3"`
	Username string `json:"username" validate:"required,gte=3"`
	Surname  string `json:"surname" validate:"required,gte=3"`
}
