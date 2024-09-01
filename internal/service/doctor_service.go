package service

import (
	"BackendTugasAkhir/entities"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type DoctorService struct {
	DB *sql.DB
}

func NewDoctorService(db *sql.DB) DoctorService {
	return DoctorService{DB: db}
}

func (h *DoctorService) AddDoctor(doctor entities.Doctor) (string, error) {
	var doctorID string

	err := h.DB.QueryRow("INSERT INTO doctors (name, hospital_id) VALUES ($1, $2) RETURNING doctor_id",
		doctor.Name, doctor.HospitalId).Scan(&doctorID)
	if err != nil {
		return "", err
	}

	return doctorID, nil
}

func (h *DoctorService) AddDoctorSchedule(doctorId string, day string, hour string) error {
	fmt.Println(hour)
	_, err := h.DB.Exec("INSERT INTO doctor_schedules (day, hour ,doctor_id) VALUES ($1, $2, $3)",
		day, hour, doctorId)
	if err != nil {
		return err
	}
	return nil
}

func (h *DoctorService) GetDoctorForAppointment(hospitalID string) ([]entities.Doctor, error) {
	var doctors []entities.Doctor
	rows, err := h.DB.Query("SELECT DISTINCT doctors.doctor_id , doctors.name, COALESCE(doctors.image, '') AS image FROM doctors JOIN doctor_schedules ON doctors.doctor_id = doctor_schedules.doctor_id WHERE doctors.hospital_id = $1", hospitalID)
	if err != nil {
		return []entities.Doctor{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var doctor entities.Doctor
		err = rows.Scan(&doctor.DoctorId, &doctor.Name, &doctor.Image)
		if err != nil {
			return []entities.Doctor{}, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (h *DoctorService) GetDoctorScheduleForAppointment(doctorId string, dateString string) ([]entities.DoctorSchedule, error) {
	var schedules []entities.DoctorSchedule
	// Parse the date string into a time.Time object
	layout := "2006-01-02"
	timestampLayout := "2006-01-02 15:04:05"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	day := date.Weekday().String()
	timestampStart := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	timestampEnd := timestampStart.AddDate(0, 0, 1)
	rows, err := h.DB.Query("SELECT ds.schedule_id, ds.hour::TEXT FROM doctor_schedules ds LEFT JOIN doctor_appointments da ON ds.schedule_id = da.schedule_id WHERE ds.day = $1 AND ds.doctor_id = $2 AND (da.schedule_id IS NULL OR da.appointment_time NOT BETWEEN $3 AND $4) ORDER BY ds.hour",
		day, doctorId, timestampStart.Format(timestampLayout), timestampEnd.Format(timestampLayout))

	if err != nil {
		fmt.Println(err)
		return []entities.DoctorSchedule{}, err
	}

	defer rows.Close()
	for rows.Next() {
		fmt.Println("rows : ", rows)
		var schedule entities.DoctorSchedule
		err = rows.Scan(&schedule.ScheduleId, &schedule.Hour)
		if err != nil {
			fmt.Println(err)
			return []entities.DoctorSchedule{}, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (h *DoctorService) GetScheduleDetail(scheduleId string) (entities.DoctorSchedule, error) {
	var schedule entities.DoctorSchedule
	rows, err := h.DB.Query("SELECT doctor_id, hour::TEXT, day FROM doctor_schedules WHERE schedule_id=$1", scheduleId)
	if err != nil {
		return entities.DoctorSchedule{}, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&schedule.DoctorId, &schedule.Hour, &schedule.Day); err != nil {
			return entities.DoctorSchedule{}, err
		}
	}
	return schedule, nil
}

func (h *DoctorService) AddDoctorOnDutySchedule(doctorOnDuty entities.DoctorOnDuty) error {
	_, err := h.DB.Exec("INSERT INTO doctor_on_duty_schedules (day, on_duty_start, on_duty_end, patient_limit, doctor_id) VALUES ($1, $2, $3, $4, $5)",
		doctorOnDuty.Day, doctorOnDuty.StartHour, doctorOnDuty.EndHour, doctorOnDuty.PatientLimit, doctorOnDuty.DoctorId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (h *DoctorService) GetDoctorOnDuty(hospitalId string) ([]entities.DoctorOnDuty, error) {
	doctors := []entities.DoctorOnDuty{}
	day := time.Now().Weekday().String()
	rows, err := h.DB.Query("SELECT ds.on_duty_id, doctors.doctor_id, doctors.name, COALESCE(doctors.image,'') AS image, ds.on_duty_start::TEXT, ds.on_duty_end::TEXT FROM doctors JOIN doctor_on_duty_schedules ds ON doctors.doctor_id = ds.doctor_id WHERE doctors.hospital_id = $1 AND ds.day = $2 ",
		hospitalId, day)

	if err != nil {
		fmt.Println(err)
		return []entities.DoctorOnDuty{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var doctor entities.DoctorOnDuty
		err = rows.Scan(&doctor.OnDutyId, &doctor.DoctorId, &doctor.Name, &doctor.Image, &doctor.StartHour, &doctor.EndHour)
		if err != nil {
			fmt.Println(err)
			return []entities.DoctorOnDuty{}, err
		}
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (h *DoctorService) GetDoctorOnDutyScheduleDetail(onDutyId string) (entities.DoctorOnDuty, error) {
	var doctor entities.DoctorOnDuty
	row := h.DB.QueryRow("SELECT day, on_duty_start::TEXT, on_duty_end::TEXT,patient_limit, doctor_id FROM doctor_on_duty_schedules WHERE on_duty_id = $1", onDutyId)
	err := row.Scan(&doctor.Day, &doctor.StartHour, &doctor.EndHour, &doctor.PatientLimit, &doctor.DoctorId)
	if err != nil {
		return entities.DoctorOnDuty{}, err
	}
	return doctor, nil
}

func (h *DoctorService) GetDoctorOnDutyLastPatient(onDutyId string) (int, string, error) {
	var queueNumber int
	var predictionTime string
	row := h.DB.QueryRow("SELECT queue_number, prediction_time::TEXT FROM patient_queues WHERE on_duty_id = $1 ORDER BY queue_number DESC ", onDutyId)
	err := row.Scan(&queueNumber, &predictionTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", nil
		} else {
			return 0, "", err
		}

	}
	return queueNumber, predictionTime, nil
}

func (h *DoctorService) AddToQueueDoctorOnDuty(queue entities.Queue) error {
	_, err := h.DB.Exec("INSERT INTO patient_queues (prediction_time, queue_number, patient_id, doctor_id, on_duty_id) VALUES ($1, $2, $3, $4, $5)",
		queue.PredictionTime, queue.QueueNumber, queue.PatientId, queue.DoctorId, queue.OnDutyId)
	if err != nil {
		return err
	}
	return nil
}
