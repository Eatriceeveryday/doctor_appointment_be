package service

import (
	"BackendTugasAkhir/entities"
	"database/sql"
)

type HospitalsService struct {
	DB *sql.DB
}

func NewHospitalsService(db *sql.DB) HospitalsService {
	return HospitalsService{DB: db}
}

func (h *HospitalsService) AddHospital(hospital entities.Hospital) (string, error) {
	var hospitalId string
	err := h.DB.QueryRow("INSERT INTO hospitals (name, address, contact_number) VALUES ($1, $2, $3) RETURNING hospital_id",
		hospital.Name, hospital.Address, hospital.ContactNumber).Scan(&hospitalId)
	if err != nil {
		return "", err
	}
	return hospitalId, nil
}

func (h *HospitalsService) GetHospital() ([]entities.Hospital, error) {
	var hospitals []entities.Hospital
	rows, err := h.DB.Query("SELECT hospital_id, name, address, contact_number, COALESCE(image, '') AS image FROM hospitals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var hospital entities.Hospital
		err := rows.Scan(&hospital.HospitalId, &hospital.Name, &hospital.Address, &hospital.ContactNumber, &hospital.Image)
		if err != nil {
			return nil, err
		}
		hospitals = append(hospitals, hospital)
	}
	return hospitals, nil
}
