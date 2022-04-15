package Dtos

//TODO: I am using this for the DB too, it should be changed but this is a practice task and I'm on holidays
type Album struct {
	ID     int64  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
