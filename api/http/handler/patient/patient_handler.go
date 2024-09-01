package patient

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/entities"
	"BackendTugasAkhir/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type PatientHandler struct {
	patientService service.PatientService
}

func NewPatientHandler(ps service.PatientService) PatientHandler {
	return PatientHandler{
		patientService: ps,
	}
}

func (h *PatientHandler) AddPatient(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)

	var req AddPatientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.patientService.AddPatient(entities.Patients{
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		UserId:      userId,
	})

	if err != nil {
		utils.JSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Patient Succesfuly Added"}, http.StatusCreated)

}

func (h *PatientHandler) GetAllPatients(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)

	patients, err := h.patientService.GetAllPatient(userId)

	if err != nil {
		utils.JSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: struct {
		Patients []entities.Patients `json:"patients"`
	}{
		Patients: patients,
	}}, http.StatusOK)
}

func (h *PatientHandler) EditPatient(w http.ResponseWriter, r *http.Request) {
	patientId := r.URL.Query().Get("patient_id")
	if patientId == "" {
		utils.JSONResponse(w, "patient_id can't be empty", http.StatusBadRequest)
		return
	}
	var req AddPatientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.patientService.EditPatient(entities.Patients{
		PatientId:   patientId,
		Name:        req.Name,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
	})
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, utils.Response{Msg: "Success"}, http.StatusAccepted)
}

func (h *PatientHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	patientId := r.URL.Query().Get("patient_id")
	if patientId == "" {
		utils.JSONResponse(w, "patient_id can't be empty", http.StatusBadRequest)
		return
	}
	err := h.patientService.DeletePatient(patientId)
	if err != nil {
		utils.JSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println(err.Error())
	}
	utils.JSONResponse(w, utils.Response{Msg: "Success delete patient"}, http.StatusOK)
}
