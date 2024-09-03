package service

import (
	"BackendTugasAkhir/entities"
	"database/sql"
)

type PatientService struct {
	DB *sql.DB
}

func NewPatientService(db *sql.DB) PatientService {
	return PatientService{DB: db}
}

func (ps *PatientService) AddPatient(patient entities.Patients) error {
	_, err := ps.DB.Exec("INSERT INTO patients (name,date_of_birth,gender, user_id) values ($1 , $2, $3, $4)",
		patient.Name, patient.DateOfBirth, patient.Gender, patient.UserId)

	if err != nil {
		return err
	}
	return nil
}

func (ps *PatientService) GetAllPatient(userId string) ([]entities.Patients, error) {
	patients := []entities.Patients{}
	rows, err := ps.DB.Query("SELECT patient_id, name, date_of_birth::TEXT, gender FROM patients WHERE user_id = $1", userId)
	if err != nil {
		return []entities.Patients{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient entities.Patients

		err := rows.Scan(&patient.PatientId, &patient.Name, &patient.DateOfBirth, &patient.Gender)
		if err != nil {
			return []entities.Patients{}, err
		}
		patients = append(patients, patient)
	}
	return patients, nil
}

func (ps *PatientService) EditPatient(patient entities.Patients) error {
	_, err := ps.DB.Exec("UPDATE patients SET name = $1, date_of_birth = $2, gender = $3 WHERE patient_id= $4",
		patient.Name, patient.DateOfBirth, patient.Gender, patient.PatientId)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PatientService) DeletePatient(patientId string) error {
	_, err := ps.DB.Exec("DELETE FROM patients WHERE patient_id = $1", patientId)
	if err != nil {
		return err
	}
	return nil
}
