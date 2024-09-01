package entities

type Doctor struct {
	DoctorId   string `json:"doctorId"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	HospitalId string `json:"hospitalId,omitempty"`
}
