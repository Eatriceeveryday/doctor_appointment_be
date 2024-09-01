package entities

type Patients struct {
	PatientId   string `json:"patientId"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	UserId      string `json:"user_id,omitempty"`
}
