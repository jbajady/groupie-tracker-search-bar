package Func

// Fetch All Data  From  Urls
func GenriateData() bool {
	er := FetchData(API.Artistes, &Artists)
	er1 := FetchData(API.Locations, &Relations)
	er2 := FetchData(API.Dates, &Relations)
	er3 := FetchData(API.Relation, &Relations)
	if er != nil || er1 != nil || er2 != nil || er3 != nil {
		return false
	}
	for i, v := range Relations.Index {
		Artists[i].Date = v.Date
		Artists[i].Location = v.Location
		Artists[i].Relation = v.Relation
	}
	return true
}
