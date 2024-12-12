package handlers

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"fmt"
	"net/http"
)

func ModulesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			err := json.NewEncoder(w).Encode(variables.Modules)
			if err != nil {
				_ = fmt.Errorf(ErrorDescription.InvalidEncode, err)
			}
			w.Header().Set(variables.ContentType, variables.ApplicationJSON)
		}
	case http.MethodPost:
		{
			var module structures.Module
			err := json.NewDecoder(r.Body).Decode(&module)
			if err != nil {
				_ = fmt.Errorf(ErrorDescription.InvalidDecode, err)
			}
			module.ID = variables.CountModulesID
			variables.CountModulesID++
			variables.Modules = append(variables.Modules, module)
			w.Header().Set(variables.ContentType, variables.ApplicationJSON)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(module)
		}
	default:
		http.Error(w, ErrorDescription.InvalidMethod, http.StatusMethodNotAllowed)
	}
}
