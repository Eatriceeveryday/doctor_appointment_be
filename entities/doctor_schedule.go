package entities

type DoctorSchedule struct {
	ScheduleId string `json:"scheduleId,omitempty"`
	Day        string `json:"day,omitempty"`
	Hour       string `json:"hour"`
	DoctorId   string `json:"doctorId,omitempty"`
}
