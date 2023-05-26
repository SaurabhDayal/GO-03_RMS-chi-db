package handlers

import (
	"06_RMS-chi-db/database/dbHelper"
	"06_RMS-chi-db/errorHandling"
	"06_RMS-chi-db/models"
	"encoding/json"
	"net/http"
)

func CreateSubAdmin(w http.ResponseWriter, r *http.Request) {
	var subAdmin models.Users
	err := json.NewDecoder(r.Body).Decode(&subAdmin)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
	}
	user, err := dbHelper.CreateNewSubAdmin(&subAdmin)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetSubAdminList(w http.ResponseWriter, r *http.Request) {
	subAdmins, err := dbHelper.GetAllSubAdmin()
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(subAdmins)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}
