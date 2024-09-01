package doctor

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/internal/service"
	"net/http"
)

type DoctorHandler struct {
	doctorService service.DoctorService
}

func NewDoctorHandler(doctorService service.DoctorService) DoctorHandler {
	return DoctorHandler{doctorService: doctorService}
}

func (h *DoctorHandler) GetDoctorAppointmentSchedule(w http.ResponseWriter, r *http.Request) {
	doctorId := r.PathValue("doctor_id")
	date := r.URL.Query().Get("date")

	if doctorId == "" || date == "" {
		utils.JSONResponse(w, utils.Response{Msg: "Invalid Request"}, http.StatusBadRequest)
		return
	}

	schedules, err := h.doctorService.GetDoctorScheduleForAppointment(doctorId, date)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: schedules}, http.StatusOK)

}
