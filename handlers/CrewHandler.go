package handlers

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/handlers/httpUtils"
	"SatarShipRESTAPI/variables"
	"net/http"
)

func CrewHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		httpUtils.HandleGetRequest(w, variables.Crew)
	case http.MethodPost:
		httpUtils.HandlePostRequest(r, w)
	default:
		http.Error(w, ErrorDescription.InvalidMethod, http.StatusMethodNotAllowed)
	}
}

func CrewMemberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		httpUtils.HandleGetRequestById(w, r)
	case http.MethodPut:
		httpUtils.HandlePutRequest(w, r)
	case http.MethodDelete:
		httpUtils.HandleDeleteMethod(w, r)
	}
}
