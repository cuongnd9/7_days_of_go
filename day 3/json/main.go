package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name, Color string
	Age         int
	Weight      float64
	IsMale      bool
}

func main() {
	chihao := Cat{
		"chihao",
		"yellow & white",
		2,
		3.8,
		true,
	}
	lily := Cat{
		"lily",
		"black",
		2,
		3.5,
		false,
	}

	// encode to JSON
	chihaoJSON, error := json.Marshal(chihao)
	if error != nil {
		fmt.Printf("error hanpended: %v", error)
		return
	}
	// print JSON string
	fmt.Println(string(chihaoJSON))
	// {"Name":"chihao","Color":"yellow \u0026 white","Age":2,"Weight":3.8,"IsMale":true}

	lilyJSON, _ := json.MarshalIndent(lily, "🐈", "   ")
	fmt.Println(string(lilyJSON))
	//{
	//🐈   "Name": "lily",
	//🐈   "Color": "black",
	//🐈   "Age": 2,
	//🐈   "Weight": 3.5,
	//🐈   "IsMale": false
	//🐈}

	jsonData := []byte(`{"name":"chiha","age":2,"color":"blue"}`)
	isValid := json.Valid(jsonData)
	fmt.Println("💅", isValid)

	var chiha Cat
	err := json.Unmarshal(jsonData, &chihao)
	fmt.Printf("😻 error: %v\n", err)
	fmt.Printf("😸: %#v", chiha)
}
