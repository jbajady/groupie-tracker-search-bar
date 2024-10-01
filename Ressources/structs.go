package Func

var (
	Artists      []Artest
	Relation     LocationData
	Artiste      Artest
	SearchArtist []Artest
	API          APi
)

type APi struct {
	Artistes  string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artest struct {
	ID            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	CreationDate  int      `json:"creationDate"`
	FirstAlbum    string   `json:"firstAlbum"`
	Date          []string
	Location      []string
	DatesLocation map[string][]string
}
type LocationData struct {
	Index []struct {
		ID            int                 `json:"id"`
		Location      []string            `json:"locations"`
		Date          []string            `json:"dates"`
		DatesLocation map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
