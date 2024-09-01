package service

import (
	"BackendTugasAkhir/entities"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type AppointmentService struct {
	DB *sql.DB
}

func NewAppointmentService(db *sql.DB) AppointmentService {
	return AppointmentService{DB: db}
}

func (h *AppointmentService) AddAppointment(appointment entities.Appointment) (string, error) {
	var appointment_id string
	err := h.DB.QueryRow("INSERT INTO doctor_appointments (appointment_time,  patient_id, schedule_id, doctor_id) VALUES ($1, $2, $3, $4) RETURNING appointment_id",
		appointment.AppointmentTime, appointment.PatientId, appointment.ScheduleId, appointment.DoctorId).Scan(&appointment_id)
	if err != nil {
		return "", err
	}

	return appointment_id, nil
}

func (h *AppointmentService) DeleteAppointment(appointment_id string) error {
	deleteAppointmentId := ""
	err := h.DB.QueryRow("DELETE FROM doctor_appointments WHERE appointment_id = $1 RETURNING appointment_id", appointment_id).Scan(&deleteAppointmentId)
	if err != nil {
		return err
	}
	if deleteAppointmentId == "" {
		return errors.New("Appointment does not exist")
	}
	return nil
}

func (h *AppointmentService) GetAppointment(patients []entities.Patients) ([]entities.Appointment, error) {
	appointments := []entities.Appointment{}
	patientIds := make([]string, len(patients))
	for i, patient := range patients {
		patientIds[i] = patient.PatientId
	}
	placeholders := make([]string, len(patientIds))
	for i := range patientIds {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	// Join the placeholders with commas
	query := fmt.Sprintf("SELECT ds.appointment_id, ds.appointment_time::TEXT, doctors.name AS doctor_name, patients.name AS patient_name FROM doctor_appointments ds JOIN doctors ON ds.doctor_id = doctors.doctor_id JOIN patients ON ds.patient_id = patients.patient_id WHERE ds.patient_id IN(%s)", strings.Join(placeholders, ","))

	// Convert the slice to an interface slice
	args := make([]interface{}, len(patientIds))
	for i, v := range patientIds {
		args[i] = v
	}

	rows, err := h.DB.Query(query, args...)

	defer rows.Close()
	for rows.Next() {
		var appointment entities.Appointment
		err = rows.Scan(&appointment.AppointmentId, &appointment.AppointmentTime, &appointment.DoctorName, &appointment.PatientName)
		if err != nil {
			return []entities.Appointment{}, err
		}
		appointments = append(appointments, appointment)
	}
	return appointments, nil
}
