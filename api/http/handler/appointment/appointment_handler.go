package appointment

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/entities"
	"BackendTugasAkhir/internal/service"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type AppointmentHandler struct {
	appointmentService service.AppointmentService
	patientService     service.PatientService
	doctorService      service.DoctorService
	validator          *validator.Validate
}

func NewAppointmentHandler(as service.AppointmentService, ps service.PatientService, v *validator.Validate, ds service.DoctorService) AppointmentHandler {
	return AppointmentHandler{
		appointmentService: as,
		patientService:     ps,
		validator:          v,
		doctorService:      ds,
	}
}

func (h *AppointmentHandler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)

	var req AppointmentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	//check if patient connected to user
	patients, err := h.patientService.GetAllPatient(userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}
	found := false
	for _, patient := range patients {
		if patient.PatientId == req.PatientId {
			found = true
			break
		}
	}

	if !found {
		utils.JSONResponse(w, utils.Response{Msg: "Invalid Patient Id Request"}, http.StatusBadRequest)
		return
	}

	schedule, err := h.doctorService.GetScheduleDetail(req.ScheduleId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	appointmentTime := req.AppointmentDate + " " + schedule.Hour

	appointmentId, err := h.appointmentService.AddAppointment(entities.Appointment{
		AppointmentTime: appointmentTime,
		DoctorId:        schedule.DoctorId,
		PatientId:       req.PatientId,
		ScheduleId:      req.ScheduleId,
	})

	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: struct {
		AppointmentId string `json:"appointmentId"`
	}{
		AppointmentId: appointmentId,
	}}, http.StatusOK)
}

func (h *AppointmentHandler) ChangeDoctorAppointment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	appointmentId := r.URL.Query().Get("appointment_id")
	if appointmentId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "appointmentId is empty"}, http.StatusBadRequest)
		return
	}

	var req AppointmentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	//check if patient connected to user
	patients, err := h.patientService.GetAllPatient(userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}
	found := false
	for _, patient := range patients {
		if patient.PatientId == req.PatientId {
			found = true
			break
		}
	}

	if !found {
		utils.JSONResponse(w, utils.Response{Msg: "Invalid Patient Id Request"}, http.StatusBadRequest)
		return
	}

	schedule, err := h.doctorService.GetScheduleDetail(req.ScheduleId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}
	layout := "2006-01-02"
	date, err := time.Parse(layout, req.AppointmentDate)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Invalid Date"}, http.StatusBadRequest)
		return
	}

	if schedule.Day != date.Weekday().String() {
		utils.JSONResponse(w, utils.Response{Msg: "Invalid Date for schedule_id"}, http.StatusBadRequest)
		return
	}

	appointmentTime := req.AppointmentDate + " " + schedule.Hour

	newAppointmentId, err := h.appointmentService.AddAppointment(entities.Appointment{
		AppointmentTime: appointmentTime,
		DoctorId:        schedule.DoctorId,
		PatientId:       req.PatientId,
		ScheduleId:      req.ScheduleId,
	})

	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	err = h.appointmentService.DeleteAppointment(appointmentId)
	if err != nil {
		//Delete newly created appointment
		newErr := h.appointmentService.DeleteAppointment(newAppointmentId)
		fmt.Println(err)
		fmt.Println(newErr)
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: struct {
		AppointmentId string `json:"appointmentId"`
	}{
		AppointmentId: newAppointmentId,
	}}, http.StatusOK)
}

func (h *AppointmentHandler) GetAppointment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	patients, err := h.patientService.GetAllPatient(userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	appointments, err := h.appointmentService.GetAppointment(patients)
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Error"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: struct {
		Appointments []entities.Appointment `json:"appointments"`
	}{
		appointments,
	}}, http.StatusOK)

}

func (h *AppointmentHandler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	appointmentId := r.URL.Query().Get("appointment_id")
	if appointmentId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "appointmentId is empty"}, http.StatusBadRequest)
		return
	}

	err := h.appointmentService.DeleteAppointment(appointmentId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Failed to delete appointment"}, http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success"}, http.StatusOK)
}
