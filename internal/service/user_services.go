package service

import (
	"BackendTugasAkhir/entities"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	DB *sql.DB
}

func NewUserServices(db *sql.DB) UserServices {
	return UserServices{DB: db}
}

func (us *UserServices) AddUser(user entities.Users) (string, error) {
	var userId string
	hashedPassword, err := us.hashPassword(user.Password)
	if err != nil {
		return "", err
	}
	err = us.DB.QueryRow("INSERT INTO users (email, username, password, contact_number) VALUES ($1, $2, $3, $4) RETURNING user_id",
		user.Email, user.UserName, hashedPassword, user.ContactNumber).Scan(&userId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return "", errors.New("User already exists")
			}
		}
		return "", err
	}
	return userId, nil
}

func (us *UserServices) GetUser(userEmail string) (entities.Users, error) {
	var user entities.Users
	row, err := us.DB.Query("SELECT user_id, email, password FROM users WHERE email = $1 ", userEmail)
	if err != nil {
		return entities.Users{}, err
	}

	for row.Next() {
		err = row.Scan(&user.UserId, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("err : " + err.Error())
			return entities.Users{}, err
		}
	}

	return user, nil

}

func (us *UserServices) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
