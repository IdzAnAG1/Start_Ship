package handlers

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/handlers/httpUtils"
	"SatarShipRESTAPI/variables"
	"fmt"
	"net/http"
)

func ModulesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		httpUtils.HandleGetRequest(w, variables.Modules)
	case http.MethodPost:
		httpUtils.HandlePostRequest(r, w)
	default:
		http.Error(w, ErrorDescription.InvalidMethod, http.StatusMethodNotAllowed)
	}
}

func ModuleByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		httpUtils.HandleGetRequestById(w, r)
	case http.MethodPut:
		fmt.Println("fdsa")
	case http.MethodDelete:
		fmt.Println("asdf")
	default:
		http.Error(w, ErrorDescription.InvalidMethod, http.StatusMethodNotAllowed)
	}
}
