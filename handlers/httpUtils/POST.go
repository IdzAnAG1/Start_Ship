package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
)

func HandlePostRequest(r *http.Request, w http.ResponseWriter) {
	ReceivedData := TypeDefinitionAndVariableDeclaration(r, w)
	err := json.NewDecoder(r.Body).Decode(ReceivedData)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidDecode, http.StatusBadRequest)
		return
	}
	structureDistributor(ReceivedData)
	w.Header().Set(variables.ContentType, variables.ApplicationJSON)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ReceivedData)
}
