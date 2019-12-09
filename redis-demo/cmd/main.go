package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type People struct {
	Name string `json:"name_title"`
	Age int `json:"age_size"`
}

func JsonToStructDemo(){
	jsonStr := `
        {
                "name_title": "jqw",
                "age_size":12
        }
        `
	var people  People
	if err := json.Unmarshal([]byte(jsonStr), &people); err != nil {
		log.Fatal(err)
	}
	fmt.Println("===========Start")
	fmt.Println(people.Name, people.Age)
	fmt.Println("===========End")
}

func main() {
	JsonToStructDemo()
}
