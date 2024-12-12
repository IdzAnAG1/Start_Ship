package handlers

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func ModuleByIDHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Path[len("/modules/"):]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidID, http.StatusBadRequest)
	}
	switch r.Method {
	case http.MethodGet:
		{
			for _, module := range variables.Modules {
				if module.ID == id {
					w.Header().Set(variables.ContentType, variables.ApplicationJSON)
					json.NewEncoder(w).Encode(module)
					return
				}
			}
			http.Error(w, ErrorDescription.ObjectNotFound, http.StatusNotFound)
		}
	case http.MethodPut:
		{
			var updatedModule structures.Module
			err := json.NewDecoder(r.Body).Decode(&updatedModule)
			if err != nil {
				http.Error(w, ErrorDescription.InvalidDecode, http.StatusBadRequest)
			}

			for i, module := range variables.Modules {
				if module.ID == id {
					variables.Modules[i].Name = updatedModule.Name
					variables.Modules[i].Status = updatedModule.Status
					variables.Modules[i].Power = updatedModule.Power
					w.Header().Set(variables.ContentType, variables.ApplicationJSON)
					err = json.NewEncoder(w).Encode(variables.Modules[i])
					if err != nil {
						http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
					}
					return
				}
			}
			http.Error(w, ErrorDescription.ObjectNotFound, http.StatusNotFound)
		}
	case http.MethodDelete:
		{
			for i, module := range variables.Modules {
				if module.ID == id {
					variables.Modules = append(variables.Modules[:i], variables.Modules[i:]...)
					w.WriteHeader(http.StatusNoContent)
				}
			}
			http.Error(w, ErrorDescription.ObjectNotFound, http.StatusNotFound)
		}
	default:
		http.Error(w, ErrorDescription.InvalidMethod, http.StatusMethodNotAllowed)
	}
}
