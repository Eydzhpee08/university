package db

import (
	"encoding/json"
	"fmt"
	jwt "github.com/Eydzhpee08/jwt"
	"github.com/Eydzhpee08/university/handlers/models"
	"github.com/Eydzhpee08/university/utils"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ResponseUser struct {
	Id    int64  `json:"id"`
	Token string `json:"token"`
}

// create employee handler
func SignIn(w http.ResponseWriter, r *http.Request) {
	var (
		login    models.Sign
		response = utils.Response{
			Code: 200,
		}
	)
	// utils.CorsOptions(w, r)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}

	auth, statusCode, err := SignIns(login.Login, login.Password)
	if err != nil {
		response.Code = statusCode
		response.Message = err.Error()
		return
	}

	response.Payload = auth
	return
}
func SignIns(login string, password string) (response ResponseUser, status int, err error) {
	var user models.Sign

	password = utils.GeneratePasswordHash(password)
	userResponse, err := GetUser(login, password)
	if err != nil {
		log.Println(err)
		return response, http.StatusNotFound, err
	}
	fmt.Println(userResponse.Login, userResponse.Name, userResponse.ID)
	var responseUser utils.Payload

	responseUser.ID = userResponse.ID
	responseUser.Exp = time.Now().Add(time.Hour * 10).Unix()
	response.Id = userResponse.ID
	response.Token, err = jwt.Encode(responseUser, jwt.Secret("SECRET"))
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}

	log.Println("Вход: Был выполнен вход пользователем { Name: " + user.Name + ", ID: " + strconv.Itoa(int(user.ID)) + "}")
	return response, http.StatusOK, nil
}

func GetUser(login string, password string) (models.Sign, error) {
	var users models.Sign

	err := Database.Table("signs").
		Where("login=? AND password=?", login, password).
		Select("id").
		First(&users).Error

	if err != nil || err == gorm.ErrRecordNotFound || password != users.Password {
		log.Println("ERROR", err)
		return users, err
	}
	return users, nil
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []models.Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}
