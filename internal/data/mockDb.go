package data

//Mock store of data. Data stores in memory
func GetDataById(id string) int {
	if id == "test" {
		return 1
	} else {
		return 100
	}
}