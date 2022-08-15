package main

import (
	structures "backend/Structures"
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cars *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conectado correctamente a mongo")

	cars = client.Database("practica1").Collection("Autos")
}

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hola mundo")
}

func Create(w http.ResponseWriter, r *http.Request) {
	newCar := structures.Car{
		Placa:  "p304kjy",
		Marca:  "Mazda",
		Modelo: 2015,
		Serie:  "3",
		Color:  "rojo",
	}

	res, err := cars.InsertOne(ctx, newCar)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auto insertado: ", res)
}

func Read(w http.ResponseWriter, r *http.Request) {
	log.Println("Lectura de datos")
}

func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("lectura de objetos")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Eliminando dato")
}

func main() {
	fmt.Println("Hola mundo")

	http.HandleFunc("/create", Create)
	http.HandleFunc("/read", Read)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	log.Println("=== backend listo")
	http.ListenAndServe(":3030", nil)
}
