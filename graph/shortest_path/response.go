package shortest_path

type Response struct {
	TotalWeight float64      //might be distance
	Polyline    [][2]float64 `json:"polyline"`
}
