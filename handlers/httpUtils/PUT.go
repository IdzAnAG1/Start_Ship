package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"SatarShipRESTAPI/structures"
	"encoding/json"
	"net/http"
	"strconv"
)

func HandlePutRequest(w http.ResponseWriter, r *http.Request) {
	path := GetRoute(r.URL.Path)
	temp := Determiner(path)
	idString := r.URL.Path[len(path):]
	id, errID := strconv.Atoi(idString)

	if errID != nil {
		http.Error(w, ErrorDescription.InvalidID, http.StatusBadRequest)
	}

	err := json.NewDecoder(r.Body).Decode(temp)

	if err != nil {
		return
	}

	switch v := temp.(type) {
	case *structures.Module:
		updateModule(w, id, *v)
		return
	case *structures.CrewMember:
		updateCrewMember(w, id, *v)
		return
	case *structures.StationResource:
		updateResource(w, id, *v)
		return
	}

	http.Error(w, ErrorDescription.ObjectNotFound, http.StatusNotFound)
}
