package errorHandling

import (
	"03_RMS/models"
	"encoding/json"
	"net/http"
)

func ErrHandle(err error, w http.ResponseWriter) {
	customErr := err.(*models.CustomClientError)
	w.WriteHeader(customErr.StatusCode)
	json.NewEncoder(w).Encode(err)
}
