package main

import(
	"github.com/keithkfield/Dino/dynowebportal"
	"os"
	"log"
	"encoding/json"
)
type configuration struct{
	Webserver string `json:"webserver"`
}
func main(){
	file,err := os.Open("config.json")
	if err != nil{
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	dynowebportal.RunWebPortal(config.Webserver)
}