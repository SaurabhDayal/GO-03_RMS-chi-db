package handlers

import (
	"06_RMS-chi-db/database/dbHelper"
	"06_RMS-chi-db/errorHandling"
	"06_RMS-chi-db/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func AddUserAddress(w http.ResponseWriter, r *http.Request) {
	var addr models.Addresses
	err := json.NewDecoder(r.Body).Decode(&addr)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadJSON(), w)
	}
	reqToken := r.Header.Get("Authorization")
	address, err := dbHelper.CreateUserNewAddress(&addr, reqToken)
	if err != nil {
		errorHandling.ErrHandle(err, w)

	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(address)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetRestaurantList(w http.ResponseWriter, r *http.Request) {
	restaurants, err := dbHelper.GetAllRestaurants()
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

func GetDishList(w http.ResponseWriter, r *http.Request) {
	strResId := chi.URLParam(r, "resId")
	resId, err := strconv.Atoi(strResId)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadURL(), w)
	}
	dishes, err := dbHelper.GetRestaurantDishes(resId)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(dishes)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func GetDistance(w http.ResponseWriter, r *http.Request) {
	strResId := chi.URLParam(r, "resId")
	resId, err := strconv.Atoi(strResId)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadURL(), w)
	}
	strUserAddId := chi.URLParam(r, "userAddId")
	userAddId, err := strconv.Atoi(strUserAddId)
	if err != nil {
		errorHandling.ErrHandle(errorHandling.UnableToReadURL(), w)
	}
	fmt.Println("resId", resId, "userAddId", userAddId)
	addDistance, err := dbHelper.GetAddDistance(resId, userAddId)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(addDistance)
		if err != nil {
			errorHandling.ErrHandle(errorHandling.UnableToWriteJSON(), w)
		}
	}
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	//var order models.Orders
	//err := json.NewDecoder(r.Body).Decode(&order)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//}
	//ord, err := dbHelper.CreateNewOrder(&order)
	//if err != nil {
	//	w.WriteHeader(http.StatusNoContent)
	//} else {
	//	w.WriteHeader(http.StatusOK)
	//	err = json.NewEncoder(w).Encode(ord)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//	}
	//}
}

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	// id distance
	//err := dbHelper.CancelOrder()
	//if err != nil {
	//	w.WriteHeader(http.StatusNoContent)
	//} else {
	//	w.WriteHeader(http.StatusOK)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//	}
	//}
}

func OkOrder(w http.ResponseWriter, r *http.Request) {
	// id distance
	//order, err := dbHelper.OkOrder()
	//if err != nil {
	//	w.WriteHeader(http.StatusNoContent)
	//} else {
	//	w.WriteHeader(http.StatusOK)
	//	err = json.NewEncoder(w).Encode(order)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//	}
	//}
}
