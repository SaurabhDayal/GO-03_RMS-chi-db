package middlewares

import (
	"06_RMS-chi-db/database/dbHelper"
	"06_RMS-chi-db/errorHandling"
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
			errorHandling.InvalidUser()
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
			errorHandling.InvalidUser()
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
			errorHandling.InvalidUser()
		}
	})
}
