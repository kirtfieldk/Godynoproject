package main

import(
	"fmt"
	"github.com/keithkfield/Dino/dynowebportal"
)

func main(){
	fmt.Println("Hello")
	dynowebportal.RunWebPortal("localhost:8787")
}