package Model

type City struct {
	Name      string  `json:"name" bson:"name"`
	Uf        string  `json:"uf" bson:"uf"`
	Latitude  float64 `json:"lat" bson:"lat"`
	Longitude float64 `json:"lon" bson:"lon"`
}
