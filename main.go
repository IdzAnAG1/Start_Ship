package main

import (
	"SatarShipRESTAPI/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Relocation)
	http.HandleFunc("/Modules", handlers.ModulesHandler)
	http.HandleFunc("/Modules/", handlers.ModuleByIDHandler)
	http.HandleFunc("/Crew", handlers.CrewHandler)
	http.HandleFunc("/Crew/", handlers.CrewMemberHandler)

	fmt.Println("Server was launched in http://localhost:8080/Modules")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
