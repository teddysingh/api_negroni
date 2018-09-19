package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/teddysingh/api_negroni/UserService/utils"

	"github.com/teddysingh/api_negroni/UserService/db"
	"github.com/teddysingh/api_negroni/UserService/models"
)

var logger = utils.Logger

// CreateUser - Handles POST action
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	logger.Printf("User is: %v", user)

	dbconn := db.DB
	if dbconn != nil {
		if err := dbconn.Create(&user).Error; err != nil {
			utils.Error(w, map[string]string{
				"Message": "Error while persisting user",
			}, http.StatusBadRequest)
		} else {
			utils.Success(w, user)
		}
	} else {
		logger.Printf("DB Connection State = %v\n", dbconn)
		utils.Error(w, map[string]string{
			"Message": "Error while connecting with DB",
		}, http.StatusInternalServerError)
	}
}

// FindUser - Get user by ID
func FindUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	logger.Printf("Form %v\n", r.URL.Query())
	queryEmail := r.URL.Query().Get("email")
	dbconn := db.DB
	if dbconn != nil {
		if err := dbconn.Where("email LIKE ?", "%"+queryEmail+"%").Find(&user).Error; err != nil {
			utils.Error(w, map[string]string{
				"Message": "Error while finding user",
			}, http.StatusNotFound)
		} else {
			utils.Success(w, user)
		}
	} else {
		logger.Printf("DB Connection State = %v\n", dbconn)
		utils.Error(w, map[string]string{
			"Message": "Error while connecting with DB",
		}, http.StatusInternalServerError)
	}
}
