package models

type User struct {
	Username string `validate:"min=3, max=40, regexp=^[a-zA-Z0-9]*$" json:"username"`
	Password string `validate:"min=3, max=40" json:"password"`
	Longname string `validate:"min=4, max=40" json:"longname"`
}
