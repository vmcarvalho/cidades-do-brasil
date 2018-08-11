package Model

type City struct {
	Name      string  `json:"name"`
	Uf        string  `json:"uf"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}
