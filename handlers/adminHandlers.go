package handlers

import (
	"03_RMS/database/dbHelper"
	"03_RMS/errorHandling"
	"03_RMS/models"
	"encoding/json"
	"net/http"
	"strconv"
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
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetSubAdminList(w http.ResponseWriter, r *http.Request) {
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		limit = 10
	}
	strOffset := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(strOffset)
	if err != nil {
		offset = 0
	}
	subAdmins, err := dbHelper.GetAllSubAdmin(limit, offset)
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
