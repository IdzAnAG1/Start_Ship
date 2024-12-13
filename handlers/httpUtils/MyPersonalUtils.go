package httpUtils

import (
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
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

//func TypeDefinitionAndVariableDeclaration(r *http.Request, w http.ResponseWriter) interface{} {
//}

func GetRoute(r *http.Request) (result string) {
	return result[:strings.Index(r.URL.Path[1:], "/")+2]
}
