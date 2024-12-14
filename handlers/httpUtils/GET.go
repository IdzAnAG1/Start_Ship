package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleGetRequest(w http.ResponseWriter, structure interface{}) {
	err := json.NewEncoder(w).Encode(structure)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
	}
	w.Header().Set(variables.ContentType, variables.ApplicationJSON)
}

func HandleGetRequestById(w http.ResponseWriter, r *http.Request) {

	path := GetRoute(r.URL.Path)

	idString := r.URL.Path[len(path):]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidID, http.StatusBadRequest)
	}

	switch path {
	case variables.ModulesRoute:
		findModuleById(id, w)
		return
	case variables.CrewRoute:
		findCrewMemberByID(id, w)
		return
	case variables.ResourceRoute:
		findStationResourceByID(id, w)
		return
	}

	http.Error(w, ErrorDescription.ObjectNotFound, http.StatusNotFound)
}
