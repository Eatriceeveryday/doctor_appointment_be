package queue

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/entities"
	"BackendTugasAkhir/internal/service"
	"fmt"
	"net/http"
	"time"
)

type QueueHandler struct {
	doctorService  service.DoctorService
	patientService service.PatientService
	queueService   service.QueueService
}

func NewHandlerQueue(doctorService service.DoctorService, ps service.PatientService, qs service.QueueService) QueueHandler {
	return QueueHandler{doctorService: doctorService, patientService: ps, queueService: qs}
}

func (hq *QueueHandler) AddQueueToDoctor(w http.ResponseWriter, r *http.Request) {
	onDutyId := r.URL.Query().Get("on_duty_id")
	if onDutyId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "Missing On Duty Id"}, http.StatusBadRequest)
		return
	}

	patient_id := r.URL.Query().Get("patient_id")
	if patient_id == "" {
		utils.JSONResponse(w, utils.Response{Msg: "Missing Patient Id"}, http.StatusBadRequest)
		return
	}

	queueDetail, err := hq.doctorService.GetDoctorOnDutyScheduleDetail(onDutyId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	if queueDetail.Day != time.Now().Weekday().String() {
		utils.JSONResponse(w, utils.Response{Msg: "Invalid On Duty Id"}, http.StatusBadRequest)
		return
	}

	startHour, err := time.Parse("15:04:05", queueDetail.StartHour)
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	endHour, err := time.Parse("15:04:05", queueDetail.EndHour)
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	now := time.Now()
	parseStart := time.Date(now.Year(), now.Month(), now.Day(), startHour.Hour(), startHour.Minute(), startHour.Second(), 0, now.Location())
	parsedEnd := time.Date(now.Year(), now.Month(), now.Day(), endHour.Hour(), endHour.Minute(), endHour.Second(), 0, now.Location())

	if now.After(parsedEnd) {
		utils.JSONResponse(w, utils.Response{Msg: "Doctor on duty time ended"}, http.StatusBadRequest)
		return
	}

	lastQueueNumber, lastPredictionTime, err := hq.doctorService.GetDoctorOnDutyLastPatient(onDutyId)
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}
	if lastQueueNumber >= queueDetail.PatientLimit {
		utils.JSONResponse(w, utils.Response{Msg: "Patient limit reached"}, http.StatusBadRequest)
		return
	}

	predictionTime := ""
	if lastQueueNumber == 0 {
		if now.After(parseStart) {
			prediction := now.Add(5 * time.Minute)
			predictionTime = prediction.Format("15:04:05")
		} else {
			predictionTime = queueDetail.StartHour
		}
	} else {
		prediction, err := time.Parse("15:04:05", lastPredictionTime)
		if err != nil {
			fmt.Println(err.Error())
			utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
			return
		}
		prediction.Add(10 * time.Minute)
		predictionTime = prediction.Format("15:04:05")
	}
	err = hq.doctorService.AddToQueueDoctorOnDuty(entities.Queue{
		PredictionTime: predictionTime,
		QueueNumber:    lastQueueNumber + 1,
		DoctorId:       queueDetail.DoctorId,
		PatientId:      patient_id,
		OnDutyId:       onDutyId,
	})
	if err != nil {
		fmt.Println("error adding queue to doctor on duty" + err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Ok"}, http.StatusCreated)
}

func (hq *QueueHandler) GetQueue(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(string)
	if userId == "" {
		utils.JSONResponse(w, utils.Response{Msg: "Missing User Id"}, http.StatusBadRequest)
		return
	}
	patients, err := hq.patientService.GetAllPatient(userId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	queues, err := hq.queueService.GetQueue(patients)
	if err != nil {
		fmt.Println(err.Error())
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, utils.Response{Msg: "Ok", Data: struct {
		Queues []entities.Queue `json:"queues"`
	}{queues}}, http.StatusOK)
}
