package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/laithrafid/bookstore_utils-go/errors_utils"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors_utils.RestErr) {
	RespondJson(w, err.Status(), err)
}
