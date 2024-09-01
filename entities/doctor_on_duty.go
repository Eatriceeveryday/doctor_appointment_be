package entities

type DoctorOnDuty struct {
	OnDutyId     string `json:"OnDutyId,omitempty"`
	Day          string `json:"day,omitempty"`
	StartHour    string `json:"startHour"`
	EndHour      string `json:"endHour"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	PatientLimit int    `json:"patientLimit,omitempty"`
	DoctorId     string `json:"doctorId,omitempty"`
}
