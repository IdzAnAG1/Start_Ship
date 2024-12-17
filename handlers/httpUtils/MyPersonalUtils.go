package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"SatarShipRESTAPI/variables"
	"encoding/json"
	"net/http"
	"sort"
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

// GetRoute - Функция принимающая на вход строку в которой должен содержаться
// URL путь, после чего выделяет значимую для определения структуры часть пути
func GetRoute(Path string) string {
	return Path[:strings.Index(Path[1:], "/")+2]
}

// Determiner - Функция распределяющая и обозначающая тип
// для дальнейшей обработки запроса
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

// findModuleById - Функция определяющая какой именно выдать результат
// по id который так же определяется исходя из пути запроса в случае,
// если Determiner определил что пришедший тип это Module из пакета
// structs
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

// findCrewMemberByID - Функция определяющая какой именно выдать результат
// по id который так же определяется исходя из пути запроса в случае,
// если Determiner определил что пришедший тип это CrewMember из пакета
// structs
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

// findStationResourceByID - Функция определяющая какой именно выдать результат
// по id который так же определяется исходя из пути запроса в случае,
// если Determiner определил что пришедший тип это StationResource из пакета
// structs
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

/*
Обдумать и переписать по возможности
*/
func deleteByID(path string, id int) {
	switch path {
	case variables.ModulesRoute:
		variables.Modules[id-1] = variables.Modules[len(variables.Modules)-1]
		variables.Modules = variables.Modules[:len(variables.Modules)-1]
		sort.Slice(variables.Modules, func(i, j int) bool {
			return variables.Modules[i].ID < variables.Modules[j].ID
		})
	case variables.CrewRoute:
		variables.Crew[id-1] = variables.Crew[len(variables.Crew)-1]
		variables.Crew = variables.Crew[:len(variables.Crew)-1]
		sort.Slice(variables.Crew, func(i, j int) bool {
			return variables.Crew[i].ID < variables.Crew[j].ID
		})
	case variables.ResourceRoute:
		variables.StationResources[id-1] = variables.StationResources[len(variables.StationResources)-1]
		variables.StationResources = variables.StationResources[:len(variables.StationResources)-1]
		sort.Slice(variables.StationResources, func(i, j int) bool {
			return variables.StationResources[i].ID < variables.StationResources[j].ID
		})
	}
}

//
//func temp [t types.Slice](slid t, i int){
//	sort.Slice(slid,func(i, j int) bool {
//		return slid[i] > slid[j]
//	})
//}
