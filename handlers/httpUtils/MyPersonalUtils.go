package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
	"strings"
)

func structureDistributor(value interface{}) {
	switch v := value.(type) {
	case *structures.Module:
		{
			v.ID = variables.CountModulesID
			variables.CountModulesID++
			variables.Modules = append(variables.Modules, *v)
		}
	case *structures.CrewMember:
		{
			v.ID = variables.CountCrewMemberID
			variables.CountCrewMemberID++
			variables.Crew = append(variables.Crew, *v)
		}
	case *structures.StationResource:
		{
			v.ID = variables.CountStationResourceID
			variables.CountStationResourceID++
			variables.StationResources = append(variables.StationResources, *v)
		}
	}
}

func GetRoute(Path string) string {
	return Path[:strings.Index(Path[1:], "/")+2]
}

func Determiner(string2 string) interface{} {
	switch string2 {
	case variables.CrewRoute:
		return &structures.CrewMember{}
	case variables.ModulesRoute:
		return &structures.Module{}
	case variables.ResourceRoute:
		return &structures.StationResource{}
	}
	return nil
}

func findModuleById(id int, w http.ResponseWriter) {
	for _, module := range variables.Modules {
		if module.ID == id {
			w.Header().Set(variables.ContentType, variables.ApplicationJSON)
			err := json.NewEncoder(w).Encode(module)
			if err != nil {
				http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
				return
			}
		}
	}
}

func findCrewMemberByID(id int, w http.ResponseWriter) {
	for _, crewMember := range variables.Crew {
		if crewMember.ID == id {
			w.Header().Set(variables.ContentType, variables.ApplicationJSON)
			err := json.NewEncoder(w).Encode(crewMember)
			if err != nil {
				http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
				return
			}
		}
	}
}

func findStationResourceByID(id int, w http.ResponseWriter) {
	for _, resource := range variables.StationResources {
		if resource.ID == id {
			w.Header().Set(variables.ContentType, variables.ApplicationJSON)
			err := json.NewEncoder(w).Encode(resource)
			if err != nil {
				http.Error(w, ErrorDescription.InvalidEncode, http.StatusBadRequest)
				return
			}
		}
	}
}
