package entities

type Queue struct {
	QueueId        string `json:"queueId"`
	PredictionTime string `json:"predictionTime"`
	QueueNumber    int    `json:"queueNumber"`
	DoctorId       string `json:"doctorId,omitempty"`
	DoctorName     string `json:"doctorName"`
	PatientId      string `json:"patientId,omitempty"`
	PatientName    string `json:"patientName"`
	OnDutyId       string `json:"onDutyId,omitempty"`
}
