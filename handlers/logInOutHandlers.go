package handlers

import (
	"06_RMS-chi-db/database/dbHelper"
	"06_RMS-chi-db/errorHandling"
	"06_RMS-chi-db/models"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var usr models.Users
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
	} else {
		user, err, token := dbHelper.RegisterUser(&usr)
		if err != nil {
			errorHandling.ErrHandle(err, w)
		} else {
			if token != "" {
				w.Header().Set("Auth-Token", token)
			}
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(user)
			if err != nil {
				errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
			}
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var usr models.Users
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
		return
	}
	var auth *models.Auths
	auth, err = dbHelper.LoginUser(usr.UserEmail, usr.UserPassword)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(auth)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("userToken")
	err := dbHelper.LogoutUser(token.(string))
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
