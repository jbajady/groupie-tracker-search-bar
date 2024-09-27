package Func

var (
	Artists  []Artest
	Relation []RelationST
	Location []Locations
	Date     []Dates
	API      APi
)

type APi struct {
	Artistes  string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artest struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations map[string][]string  `json:"locations"`
}
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type RelationST struct {
	DatesLocation map[string][]string `json:"datesLocations"`
}
type DataFinal struct {
	Artiste  Artest
	Relation RelationST
	Location Locations
	Date     Dates
}
