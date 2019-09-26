package main

import (
	"fmt"
	"log"
	"os"

	"github.com/keithkfield/Dino/dynowebportal"

	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type configuration struct {
	Webserver string `json:"webserver"`
	MongoKey  string `json: "mongoKey"`
}
type animal struct {
	// ID         primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	AnimalType string `json:"animalType" bson:"animalType"`
	Zone       int    `json:"zone" bson:"zone"`
	Nickname   string `json:"nickname" bson:"nickname"`
	Age        int    `json:"age" bson:"age"`
}

// func getAnimals(response http.ResponseWritter, request *http.Request) {}
func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoKey)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	collection := client.Database("Dino").Collection("animals")
	// Insert Docs
	// animals := []interface{}{animal{
	// 	AnimalType: "Dino",
	// 	Zone:       2,
	// 	Nickname:   "Joey",
	// 	Age:        2,
	// }, animal{
	// 	AnimalType: "Dino",
	// 	Zone:       1,
	// 	Nickname:   "Dillion",
	// 	Age:        2,
	// }, animal{
	// 	AnimalType: "cat",
	// 	Zone:       3,
	// 	Nickname:   "Billy",
	// 	Age:        22,
	// }}
	// newAnimal := []interface{}{animal{
	// 	AnimalType: "Pig",
	// 	Zone:       5,
	// 	Nickname:   "Piggie",
	// 	Age:        22,
	// }}
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = collection.InsertMany(context.TODO(), animals)
	// _, err = collection.InsertOne(context.TODO(), newAnimal[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// To delete a secific element with flter
	filter := bson.M{"nickname": "Joey"}
	collection.DeleteOne(context.Background(), filter)
	filterTwo := bson.M{"nickname": "Billy"}
	collection.DeleteMany(context.Background(), filterTwo)
	// Updating the database
	filterThree := bson.M{"nickname": "Piggie"}
	collection.UpdateOne(context.Background(), filterThree, bson.D{
		{"$set", bson.D{{"age", 33}}}})

	// fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	fmt.Println("Connected to MongoDB!")

	// router := mux.NewRouter()
	// router.HandleFunc("/api/animals", getAnimals).Methods("GET")

	dynowebportal.RunWebPortal(config.Webserver)

}
