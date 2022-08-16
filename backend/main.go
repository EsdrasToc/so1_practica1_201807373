package main

import (
	structures "backend/Structures"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

func Create(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	newCar := structures.Car{}

	json.Unmarshal(body, &newCar)

	res, err := cars.InsertOne(ctx, newCar)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auto insertado: ", res)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	findOpts := options.Find()

	var results []structures.Car
	response := ""

	cur, err := cars.Find(context.TODO(), bson.D{{}}, findOpts)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var s structures.Car
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, s)

		j, _ := json.MarshalIndent(results[len(results)-1], "", "\t")
		response = response + string(j) + ","
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	response = "[" + response[:len(response)-1] + "]"

	fmt.Fprintln(w, response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("lectura de objetos")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Eliminando dato")
}

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/create", Create).Methods("POST")
	r.HandleFunc("/readall", ReadAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func main() {
	s := New()
	log.Fatal(http.ListenAndServe(":3000", s.Router()))
}
