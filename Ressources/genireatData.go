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
	return true
}
