package structures

type Module struct {
	ID     int    `json:"ID"`
	Name   string `json:"Name"`
	Status string `json:"Status"`
	Power  int    `json:"Power"`
}

type CrewMember struct {
	ID              int    `json:"ID"`
	Name            string `json:"Name"`
	CurrentPosition string `json:"CurrentPosition"`
	ModuleID        string `json:"ModuleID"`
}

type StationResource struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Amount   int    `json:"Amount"`
	Critical string `json:"Critical"`
}
