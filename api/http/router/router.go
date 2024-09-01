package router

import (
	"BackendTugasAkhir/api/http/handler/appointment"
	"BackendTugasAkhir/api/http/handler/doctor"
	"BackendTugasAkhir/api/http/handler/hospital"
	"BackendTugasAkhir/api/http/handler/login"
	"BackendTugasAkhir/api/http/handler/patient"
	"BackendTugasAkhir/api/http/handler/queue"
	"BackendTugasAkhir/api/http/handler/register"
	"BackendTugasAkhir/api/http/middleware"
	"net/http"
)

func CreateNewRouter(rh register.Handler, lh login.LoginHandler, ph patient.PatientHandler, hh hospital.HospitalHandler, dh doctor.DoctorHandler, ah appointment.AppointmentHandler, qh queue.QueueHandler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /register", rh.Register)
	router.HandleFunc("POST /login", lh.Login)

	ProtectedRoute := http.NewServeMux()
	ProtectedRoute.HandleFunc("POST /patient", ph.AddPatient)
	ProtectedRoute.HandleFunc("GET /patient", ph.GetAllPatients)
	ProtectedRoute.HandleFunc("PUT /patient", ph.EditPatient)
	ProtectedRoute.HandleFunc("DELETE /patient", ph.DeletePatient)
	ProtectedRoute.HandleFunc("GET /hospital", hh.GetHospitals)
	ProtectedRoute.HandleFunc("GET /hospital/appointment", hh.GetDoctorWithAppointment)
	ProtectedRoute.HandleFunc("GET /hospital/appointment/{doctor_id}", dh.GetDoctorAppointmentSchedule)
	ProtectedRoute.HandleFunc("POST /appointment", ah.CreateAppointment)
	ProtectedRoute.HandleFunc("PUT /appointment", ah.ChangeDoctorAppointment)
	ProtectedRoute.HandleFunc("GET /appointment", ah.GetAppointment)
	ProtectedRoute.HandleFunc("DELETE /appointment", ah.DeleteAppointment)
	ProtectedRoute.HandleFunc("GET /hospital/on-duty", hh.GetDoctorOnDuty)
	ProtectedRoute.HandleFunc("POST /queue", qh.AddQueueToDoctor)
	ProtectedRoute.HandleFunc("GET /queue", qh.GetQueue)
	router.Handle("/", middleware.AuthenticateToken(ProtectedRoute))

	return router
}
