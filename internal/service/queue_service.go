package service

import (
	"BackendTugasAkhir/entities"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type QueueService struct {
	DB *sql.DB
}

func NewQueueService(db *sql.DB) QueueService {
	return QueueService{DB: db}
}

func (h *QueueService) GetQueue(patients []entities.Patients) ([]entities.Queue, error) {
	queues := []entities.Queue{}
	patientIds := make([]string, len(patients))
	for i, patient := range patients {
		patientIds[i] = patient.PatientId
	}
	placeholders := make([]string, len(patientIds))
	for i := range patientIds {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf("SELECT pq.queue_id, pq.prediction_time::TEXT, pq.queue_number ,doctors.name AS doctor_name, patients.name AS patient_name FROM patient_queues pq JOIN doctors ON pq.doctor_id = doctors.doctor_id JOIN patients ON pq.patient_id = patients.patient_id WHERE pq.patient_id IN(%s)", strings.Join(placeholders, ","))

	args := make([]interface{}, len(patientIds))
	for i, v := range patientIds {
		args[i] = v
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return queues, nil
		}
		return queues, err
	}
	defer rows.Close()
	for rows.Next() {
		var queue entities.Queue
		err = rows.Scan(&queue.QueueId, &queue.PredictionTime, &queue.QueueNumber, &queue.DoctorName, &queue.PatientName)
		if err != nil {
			return queues, err
		}

		queues = append(queues, queue)
	}
	return queues, nil
}
