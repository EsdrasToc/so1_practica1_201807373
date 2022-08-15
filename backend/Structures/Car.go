package structures

//import "bson"

type Car struct {
	Placa  string `bson:"placa" json:"placa"`
	Marca  string `bson:"marca" json:"marca"`
	Modelo int    `bson:"modelo" json:"modelo"`
	Serie  string `bson:"serie" json:"serie"`
	Color  string `bson:"color" json:"color"`
}
