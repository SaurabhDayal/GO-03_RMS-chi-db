package errorHandling

import (
	"06_RMS-chi-db/models"
	"encoding/json"
	"net/http"
)

func ErrHandle(err error, w http.ResponseWriter) {
	customErr := err.(*models.CustomClientError)
	w.WriteHeader(customErr.StatusCode)
	json.NewEncoder(w).Encode(err)
}
