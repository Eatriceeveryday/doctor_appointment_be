package patient

type AddPatientRequest struct {
	Name        string `json:"name" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
	Gender      string `json:"gender" validate:"required,oneof: Laki-laki Perempuan"`
}
