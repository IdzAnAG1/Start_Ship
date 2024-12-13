package main

import (
	"SatarShipRESTAPI/handlers/httpUtils"
	"SatarShipRESTAPI/structures"
	"fmt"
)

func main() {
	temp := []structures.Module{
		{
			ID:     1,
			Name:   "Temp1",
			Status: "ok",
			Power:  103,
		},
		{
			ID:     2,
			Name:   "Temp2",
			Status: "ok",
			Power:  103,
		},
	}
	temp2 := []structures.CrewMember{
		{
			ID:              1,
			Name:            "Temp1",
			CurrentPosition: "Captain",
			ModuleID:        temp[1].Name,
		},
		{
			ID:              2,
			Name:            "Temp2",
			CurrentPosition: "Captain",
			ModuleID:        temp[1].Name,
		},
	}

	test2 := httpUtils.StructInitializer(temp[0])
	test3 := httpUtils.StructInitializer(temp2[0])
	fmt.Println(test2, "- test2 |", test3, "- test3")
}
