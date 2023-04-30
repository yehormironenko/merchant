package requests

type RegisterUser struct {
	Username  string `json:"username" validate:"required,gte=3"`
	Firstname string `json:"firstname" validate:"required,gte=3"`
	Surname   string `json:"surname" validate:"required,gte=3"`
}
