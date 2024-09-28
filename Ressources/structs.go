package Func

var (
	Artists  []Artest
	Relation RelationST
	Lcation      LocationData
	Date     Dates
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
type LocationData struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type RelationST struct {
	Index []struct {
		DatesLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
type DataFinal struct {
	Artiste  Artest
	Relation RelationST
	Location LocationData
	Date     Dates
}
