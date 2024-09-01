package appointment

type AppointmentRequest struct {
	PatientId       string `json:"patientId" validate:"required"`
	ScheduleId      string `json:"scheduleId" validate:"required"`
	AppointmentDate string `json:"appointmentDate" validate:"required"`
}
