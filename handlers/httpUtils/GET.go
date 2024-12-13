package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, structure interface{}) {
	err := json.NewEncoder(w).Encode(structure)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
	}
	w.Header().Set(variables.ContentType, variables.ApplicationJSON)
}

func HandleGetRequestById(w http.ResponseWriter, r *http.Request) {
	//path := r.URL.Path
}
