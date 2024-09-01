package entities

type Appointment struct {
	AppointmentId   string `json:"appointmentId"`
	AppointmentTime string `json:"appointmentTime"`
	DoctorId        string `json:"doctorId,omitempty"`
	PatientId       string `json:"patientId,omitempty"`
	ScheduleId      string `json:"scheduleId,omitempty"`
	PatientName     string `json:"patientName,omitempty"`
	DoctorName      string `json:"doctorName,omitempty"`
}
