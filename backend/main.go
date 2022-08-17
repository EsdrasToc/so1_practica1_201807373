package main

import (
	structures "backend/Structures"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var cars, logs *mongo.Collection
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
	logs = client.Database("practica1").Collection("Logs")
}

func Create(w http.ResponseWriter, r *http.Request) {

	logs.InsertOne(ctx, structures.NewLog("Create"))

	body, _ := ioutil.ReadAll(r.Body)

	newCar := structures.Car{}

	json.Unmarshal(body, &newCar)

	res, err := cars.InsertOne(ctx, newCar)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auto insertado: ", res)
}

func ReadWithFilter(w http.ResponseWriter, r *http.Request) {
	logs.InsertOne(ctx, structures.NewLog("Read"))
	search, _ := mux.Vars(r)["find"]
	val, _ := strconv.Atoi(mux.Vars(r)["val"])

	findOpts := options.Find()

	var results []structures.Car
	response := ""

	filter := bson.D{{}}

	if val == 1 {
		filter = bson.D{{"marca", search}}
	} else if val == 2 {
		v, _ := strconv.Atoi(search)
		filter = bson.D{{"modelo", v}}
	} else {
		filter = bson.D{{"color", search}}
	}

	cur, err := cars.Find(context.TODO(), filter, findOpts)
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

func ReadAll(w http.ResponseWriter, r *http.Request) {
	logs.InsertOne(ctx, structures.NewLog("Read"))
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
	logs.InsertOne(ctx, structures.NewLog("Update"))
	body, _ := ioutil.ReadAll(r.Body)

	newCar := structures.Car{}

	json.Unmarshal(body, &newCar)

	filter := bson.D{{"placa", newCar.Placa}}

	updateResult, err := cars.UpdateOne(context.TODO(), filter, bson.D{
		{"$set", bson.D{
			{"marca", newCar.Marca},
			{"modelo", newCar.Modelo},
			{"serie", newCar.Serie},
			{"color", newCar.Color},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	logs.InsertOne(ctx, structures.NewLog("Delete"))
	body, _ := ioutil.ReadAll(r.Body)

	newCar := structures.Car{}

	json.Unmarshal(body, &newCar)

	filter := bson.D{{"placa", newCar.Placa}}

	deleteResult, err := cars.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Vehiculo removido: ", deleteResult.DeletedCount)
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
	r.HandleFunc("/filter/{val:[1-3]}/{find:.+}", ReadWithFilter).Methods("GET")
	r.HandleFunc("/delete", Delete).Methods("POST")
	r.HandleFunc("/update", Update).Methods("POST")
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
