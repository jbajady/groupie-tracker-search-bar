package Func

func GenriateData() bool {
	er := FetchData(API.Artistes, &Artists)
	er1 := FetchData(API.Locations, &Location)
	er2 := FetchData(API.Dates, &Date)
	er3 := FetchData(API.Relation, &Relation)
	if er != nil || er1 != nil || er2 != nil || er3 != nil {
		return false
	}
	return true
}
