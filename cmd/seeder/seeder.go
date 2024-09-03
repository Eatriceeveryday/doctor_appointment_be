package main

import (
	"BackendTugasAkhir/entities"
	"BackendTugasAkhir/internal/config"
	"BackendTugasAkhir/internal/database"
	"BackendTugasAkhir/internal/service"
	"fmt"
	"github.com/jaswdr/faker"
	"time"
)

var userService service.UserServices
var patientService service.PatientService

var hospitalService service.HospitalsService
var doctorService service.DoctorService
var fakerObj faker.Faker

func init() {
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	err = database.ConnectDatabase(*config.Cfg)
	if err != nil {
		panic(err)
	}

	fakerObj = faker.New()
	hospitalService = service.NewHospitalsService(database.DB)
	doctorService = service.NewDoctorService(database.DB)

}

func main() {
	fmt.Println("Seeding DB")

	for i := 0; i < 2; i++ {
		hospitalId, err := hospitalService.AddHospital(entities.Hospital{
			Name:          "Rumah Sakit " + fakerObj.Company().Name(),
			Address:       fakerObj.Address().Address(),
			ContactNumber: "+628-1234-4432-12",
			Image:         "https://mysiloam-api.siloamhospitals.com/public-asset/website-cms/website-cms-16273629729867208.jpg",
		})
		if err != nil {
			panic(err)
		}
		days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
		for j := 0; j < 5; j++ {
			doctorId, err := doctorService.AddDoctor(entities.Doctor{
				Name:       fakerObj.Person().Name(),
				HospitalId: hospitalId,
				Image:      "https://img.freepik.com/free-photo/healthcare-workers-preventing-virus-quarantine-campaign-concept-cheerful-friendly-asian-female-physician-doctor-with-clipboard-daily-checkup-standing-white-background_1258-107867.jpg?t=st=1725206635~exp=1725210235~hmac=530ce81a8db49a367f1db2d4ba1478fe11615de60264e182d8f2d2060567eb90&w=996",
			})
			if err != nil {
				panic(err)
			}

			baseTime := time.Date(2024, 8, 28, 0, 0, 0, 0, time.UTC)
			hour := baseTime.Add(8 * time.Hour)
			for k := 0; k < 5; k++ {
				for z := 0; z < 5; z++ {
					err = doctorService.AddDoctorSchedule(doctorId, days[k], hour.Format("15:04:05"))
					if err != nil {
						panic(err)
					}
					hour = hour.Add(time.Duration(30) * time.Minute)
				}
				hour = hour.Add(-time.Duration(150) * time.Minute)
			}
		}

		for x := 0; x < 3; x++ {
			doctorId, err := doctorService.AddDoctor(entities.Doctor{
				Name:       fakerObj.Person().Name(),
				HospitalId: hospitalId,
				Image:      "https://img.freepik.com/free-photo/healthcare-workers-preventing-virus-quarantine-campaign-concept-cheerful-friendly-asian-female-physician-doctor-with-clipboard-daily-checkup-standing-white-background_1258-107867.jpg?t=st=1725206635~exp=1725210235~hmac=530ce81a8db49a367f1db2d4ba1478fe11615de60264e182d8f2d2060567eb90&w=996",
			})
			if err != nil {
				panic(err)
			}
			for k := 0; k < 5; k++ {
				err = doctorService.AddDoctorOnDutySchedule(entities.DoctorOnDuty{
					Day:          days[k],
					StartHour:    "08:00",
					EndHour:      "14:00",
					PatientLimit: 20,
					DoctorId:     doctorId,
				})
				if err != nil {
					panic(err)
				}

			}
		}
	}

	fmt.Println("Seeding Done")
}
