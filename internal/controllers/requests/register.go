package requests

type RegisterUser struct {
	Username string `json:"username" binding:"required,gte=3"`
	Longname string `json:"longname" binding:"required,gte=5"`
}
