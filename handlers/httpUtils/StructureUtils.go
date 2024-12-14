package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
)

func updateModule(w http.ResponseWriter, id int, m structures.Module) {
	for i, elem := range variables.Modules {
		if elem.ID == id {
			variables.Modules[i].Name = m.Name
			variables.Modules[i].Power = m.Power
			variables.Modules[i].Status = m.Status
			err := json.NewEncoder(w).Encode(variables.Modules[i])
			if err != nil {
				http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
				return
			}
			break
		}
	}
}

func updateCrewMember(w http.ResponseWriter, id int, m structures.CrewMember) {
	for i, elem := range variables.Crew {
		if elem.ID == id {
			variables.Crew[i].Name = m.Name
			variables.Crew[i].CurrentPosition = m.CurrentPosition
			variables.Crew[i].ModuleID = m.ModuleID
			err := json.NewEncoder(w).Encode(variables.Crew[i])
			if err != nil {
				http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
				return
			}
			break
		}
	}
}

func updateResource(w http.ResponseWriter, id int, m structures.StationResource) {
	for i, elem := range variables.StationResources {
		if elem.ID == id {
			variables.StationResources[i].Name = m.Name
			variables.StationResources[i].Amount = m.Amount
			variables.StationResources[i].Critical = m.Critical
			err := json.NewEncoder(w).Encode(variables.StationResources[i])
			if err != nil {
				http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
				return
			}
			break
		}
	}
}
