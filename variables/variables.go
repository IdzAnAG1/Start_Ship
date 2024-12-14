package variables

import "SatarShipRESTAPI/structures"

var CountModulesID int = 1
var CountCrewMemberID int = 1
var CountStationResourceID int = 1
var (
	ResourceRoute    string = "/Resource/"
	CrewRoute        string = "/Crew/"
	ModulesRoute     string = "/Modules/"
	ContentType      string = "Content-Type"
	ApplicationJSON  string = "application/json"
	Modules          []structures.Module
	Crew             []structures.CrewMember
	StationResources []structures.StationResource
)
