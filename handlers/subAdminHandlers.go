package handlers

import (
	"06_RMS-chi-db/database/dbHelper"
	"06_RMS-chi-db/errorHandling"
	"06_RMS-chi-db/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var restaurant models.Restaurants
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
	}
	reqToken := r.Header.Get("Authorization")
	//token := r.Context().Value("userToken")
	resta, err := dbHelper.CreateNewRestaurant(&restaurant, reqToken)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(resta)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func CreateDish(w http.ResponseWriter, r *http.Request) {
	var dish models.Dishes
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
	}
	reqToken := r.Header.Get("Authorization")
	dishGot, err := dbHelper.CreateNewDish(&dish, reqToken)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(dishGot)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetRestaurantByOwnerId(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	restaurants, err := dbHelper.GetMyRestaurants(reqToken)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(restaurants)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetDishByOwnerId(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")
	dishes, err := dbHelper.GetMyDishes(reqToken)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(dishes)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetUsersList(w http.ResponseWriter, r *http.Request) {
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
	users, err := dbHelper.GetUsersLimitOffset(limit, offset)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func AddOwnerAddress(w http.ResponseWriter, r *http.Request) {
	var addr models.Addresses
	err := json.NewDecoder(r.Body).Decode(&addr)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
	}
	address, err := dbHelper.CreateOwnerNewAddress(addr)
	if err != nil {
		errorHandling.ErrHandle(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(address)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}
