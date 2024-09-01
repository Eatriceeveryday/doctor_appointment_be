package entities

type Hospital struct {
	HospitalId    string `json:"hospitalId"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	Address       string `json:"address"`
	ContactNumber string `json:"contactNumber"`
}
