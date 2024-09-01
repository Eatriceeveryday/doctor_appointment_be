package register

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/entities"
	"BackendTugasAkhir/internal/service"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	userService    service.UserServices
	validator      *validator.Validate
	patientService service.PatientService
}

func NewRegisterHandler(userService service.UserServices, validator *validator.Validate, patientService service.PatientService) Handler {
	return Handler{userService: userService, validator: validator, patientService: patientService}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
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

	userid, err := h.userService.AddUser(entities.Users{
		UserName:      req.Username,
		Email:         req.Email,
		Password:      req.Password,
		ContactNumber: req.ContactNumber,
	})

	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	err = h.patientService.AddPatient(entities.Patients{
		Name:        req.Username,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		UserId:      userid,
	})

	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Success"}, http.StatusCreated)
}
