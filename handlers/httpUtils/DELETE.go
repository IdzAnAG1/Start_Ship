package httpUtils

import (
	"SatarShipRESTAPI/ErrorDescription"
	"net/http"
	"strconv"
)

func HandleDeleteMethod(w http.ResponseWriter, r *http.Request) {
	path := GetRoute(r.URL.Path)
	idStr := r.URL.Path[len(path):]
	var id, err = strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrorDescription.InvalidID, http.StatusBadRequest)
	}
	deleteByID(path, id)
	w.WriteHeader(http.StatusOK)
}

/*
func sortStructuresSlice[T constraints.Ordered](val []T) {
	for i := 0; i < len(val); i++ {
		for j := i + 1; j < len(val); j++ {
			if val[i] < val[j] {
				val[i], val[j] = val[j], val[i]
			}
		}
	}

}*/
