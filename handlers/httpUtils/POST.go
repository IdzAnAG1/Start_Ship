package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
)

func HandlePostRequest(r *http.Request, w http.ResponseWriter) {
	var temp interface{}

	switch r.URL.Path {
	case "/Modules/", "/Modules":
		{
			temp = &structures.Module{}
		}
	case "/Crew/", "/Crew":
		{
			temp = &structures.CrewMember{}
		}
	case "/Resources/", "/Resources":
		{
			temp = &structures.StationResource{}
		}
	default:
		http.Error(w, ErrorDescription.UnknownEndpoint, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(temp)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidDecode, http.StatusBadRequest)
		return
	}
	structureDistributor(temp)
	w.Header().Set(variables.ContentType, variables.ApplicationJSON)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(temp)
}
