package models

type User struct {
	Username string `validate:"min=2, max=40, regexp=^[a-zA-Z0-9]*$" json:"username"`
	Surname  string `validate:"min=2, max=40, regexp=^[a-zA-Z0-9]*$" json:"surname"`
}
