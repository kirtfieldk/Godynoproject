package main

import(
	"github.com/keithkfield/Dino/dynowebportal"
	"os"
	"log"
	"fmt"
	"context"
	"encoding/json"
	// "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
type configuration struct{
	Webserver string `json:"webserver"`
}
type dino struct{
	name string
	age string
}
func main(){
	// Set client options
clientOptions := options.Client().ApplyURI("mongodb+srv://keithkfield:Icecat12!@goconnect-glenv.mongodb.net/test?retryWrites=true&w=majority")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)
// collection := client.Database("test").Collection("trainers")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")
	file,err := os.Open("config.json")
	if err != nil{
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	dynowebportal.RunWebPortal(config.Webserver)
}