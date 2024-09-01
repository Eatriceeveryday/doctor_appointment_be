package login

import (
	"BackendTugasAkhir/api/http/utils"
	"BackendTugasAkhir/internal/service"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type LoginHandler struct {
	userService    service.UserServices
	validator      *validator.Validate
	patientService service.PatientService
}

func NewLoginHandler(userService service.UserServices, validator *validator.Validate, patientService service.PatientService) LoginHandler {
	return LoginHandler{
		userService:    userService,
		validator:      validator,
		patientService: patientService,
	}
}

func (h LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	user, err := h.userService.GetUser(req.Email)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Password Salah"}, http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	token, err := createToken(user.UserId)
	if err != nil {
		utils.JSONResponse(w, utils.Response{Msg: "Internal Server Error "}, http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	fmt.Println(token)

	utils.JSONResponse(w, utils.Response{
		Msg: "Login Success",
		Data: struct {
			Token string `json:"token"`
		}{
			Token: token,
		},
	}, http.StatusOK)
}

func createToken(userId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}
