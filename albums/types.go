package albums

type Album struct {
	ID     int64   `json:"id"`
	Artist int64  `json:"artist_id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}
