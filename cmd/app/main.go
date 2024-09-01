package main

import (
	"BackendTugasAkhir/api/http"
	"BackendTugasAkhir/api/http/handler/appointment"
	"BackendTugasAkhir/api/http/handler/doctor"
	"BackendTugasAkhir/api/http/handler/hospital"
	"BackendTugasAkhir/api/http/handler/login"
	"BackendTugasAkhir/api/http/handler/patient"
	"BackendTugasAkhir/api/http/handler/queue"
	"BackendTugasAkhir/api/http/handler/register"
	httprouter "BackendTugasAkhir/api/http/router"
	"BackendTugasAkhir/internal/config"
	"BackendTugasAkhir/internal/database"
	"BackendTugasAkhir/internal/service"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	http2 "net/http"
)

var server http2.Server

func init() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	err = database.ConnectDatabase(*config.Cfg)
	if err != nil {
		panic(err)
	}
	validator := validator2.New()

	userService := service.NewUserServices(database.DB)
	patientService := service.NewPatientService(database.DB)
	hospitalService := service.NewHospitalsService(database.DB)
	doctorService := service.NewDoctorService(database.DB)
	queueService := service.NewQueueService(database.DB)
	appointmentService := service.NewAppointmentService(database.DB)
	registerHandler := register.NewRegisterHandler(userService, validator, patientService)
	loginHandler := login.NewLoginHandler(userService, validator, patientService)
	patientHandler := patient.NewPatientHandler(patientService)
	hospitalHandler := hospital.NewHospitalHandler(hospitalService, doctorService)
	doctorHandler := doctor.NewDoctorHandler(doctorService)
	appointmentHandler := appointment.NewAppointmentHandler(appointmentService, patientService, validator, doctorService)
	queueHandler := queue.NewHandlerQueue(doctorService, patientService, queueService)
	router := httprouter.CreateNewRouter(registerHandler, loginHandler, patientHandler, hospitalHandler, doctorHandler, appointmentHandler, queueHandler)
	server = http.CreateNewServer(router)
}

func main() {
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
