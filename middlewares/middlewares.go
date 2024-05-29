package middlewares

import (
	"03_RMS/database/dbHelper"
	"context"
	"net/http"
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func VerifyAdminMidd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		checkUser := dbHelper.CheckUserRole(reqToken)
		r = r.WithContext(context.WithValue(r.Context(), "userToken", reqToken))
		if checkUser == "admin" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

func VerifySubAdminMidd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		checkUser := dbHelper.CheckUserRole(reqToken)
		r = r.WithContext(context.WithValue(r.Context(), "userToken", reqToken))
		if checkUser == "subAdmin" || checkUser == "admin" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

func VerifyUserMidd(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		checkUser := dbHelper.CheckUserRole(reqToken)
		r = r.WithContext(context.WithValue(r.Context(), "userToken", reqToken))
		if checkUser == "user" || checkUser == "admin" || checkUser == "subAdmin" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
