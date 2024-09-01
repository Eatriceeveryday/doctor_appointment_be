package register

type RegisterRequest struct {
	Email         string `json:"email" validate:"required,email"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	DateOfBirth   string `json:"dateOfBirth" validate:"required"`
	Gender        string `json:"gender" validate:"required,oneof=Laki-laki Perempuan"`
	ContactNumber string `json:"contactNumber" validate:"required"`
}
