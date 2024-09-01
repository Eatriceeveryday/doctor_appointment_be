package hospital

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/entities"
	"BackendTugasAkhir/internal/service"
	"fmt"
	"net/http"
)

type HospitalHandler struct {
	hospitalService service.HospitalsService
	doctorService   service.DoctorService
}

func NewHospitalHandler(hs service.HospitalsService, ds service.DoctorService) HospitalHandler {
	return HospitalHandler{hospitalService: hs, doctorService: ds}
}

func (h *HospitalHandler) GetHospitals(w http.ResponseWriter, r *http.Request) {
	hospitals, err := h.hospitalService.GetHospital()
	if err != nil {
		fmt.Println("Error Getting hospital data, err : " + err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, utils.Response{
		Msg: "Success",
		Data: struct {
			Hospitals []entities.Hospital `json:"hospitals"`
		}{
			Hospitals: hospitals,
		},
	}, http.StatusOK)

}

func (h *HospitalHandler) GetDoctorWithAppointment(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hospitalId := query.Get("hospital_id")
	if hospitalId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "Hospital Id is missing"}, http.StatusBadRequest)
		return
	}
	doctors, err := h.doctorService.GetDoctorForAppointment(hospitalId)
	if err != nil {
		fmt.Println("Error Getting hospital data, err : " + err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, utils.Response{
		Msg: "Success",
		Data: struct {
			Doctors []entities.Doctor `json:"doctors"`
		}{
			Doctors: doctors,
		},
	}, http.StatusOK)
}

func (h *HospitalHandler) GetDoctorOnDuty(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hospitalId := query.Get("hospital_id")
	if hospitalId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "Hospital Id is missing"}, http.StatusBadRequest)
		return
	}

	doctors, err := h.doctorService.GetDoctorOnDuty(hospitalId)
	if err != nil {
		fmt.Println("Error Getting hospital data, err : " + err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w, utils.Response{Msg: "Success", Data: struct {
		Doctors []entities.DoctorOnDuty `json:"doctors"`
	}{
		doctors,
	}}, http.StatusOK)
}
