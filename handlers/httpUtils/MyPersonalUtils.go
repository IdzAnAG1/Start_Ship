package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"net/http"
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

func TypeDefinitionAndVariableDeclaration(r *http.Request, w http.ResponseWriter) interface{} {
	var temp interface{}
	switch r.URL.Path {
	case "/Modules/", "/Module":
		{
			temp = &structures.Module{}
			return temp
		}
	case "/Crew/", "/Crew":
		{
			temp = &structures.CrewMember{}
			return temp
		}
	case "/Resources/", "/Resources":
		{
			temp = &structures.StationResource{}
			return temp
		}
	default:
		http.Error(w, ErrorDescription.UnknownEndpoint, http.StatusBadRequest)
		return nil
	}
}
