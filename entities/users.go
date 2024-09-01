package entities

type Users struct {
	UserId        string `json:"userId"`
	UserName      string `json:"userName"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	ContactNumber string `json:"contactNumber"`
}
