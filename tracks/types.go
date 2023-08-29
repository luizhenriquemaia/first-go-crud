package tracks

type Track struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Order_number int16  `json:"order_number"`
	Duration     int64  `json:"duration"`
	Album        int64  `json:"album_id"`
}
