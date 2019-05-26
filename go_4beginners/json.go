package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"-"`
	Age int `json:"age"`
	Address string `json:"address,omitempty"`
	memo string
}

func main() {
	fmt.Println("Write")

	person := &Person{
		ID: 1,
		Name: "Go",
		Email: "go@mail",
		Age: 1,
		Address: "",
		memo: "test",
	}

	fmt.Println(person)

	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read")

	file2, err2 := os.Open("./person.json")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file2.Close()

	var person2 Person
	
	decoder := json.NewDecoder(file2)

	err2 = decoder.Decode(&person2)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(person2)
}
