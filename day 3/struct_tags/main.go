package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Go struct tags are annotations that appear after the type in a Go struct declaration. Each tag is composed of short strings associated with some corresponding value.

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"` // private field
	PreferredFish []string  `json:"preferredFish,omitempty"` // omitempty
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
	//{
	//	"name": "Sammy the Shark",
	//	"password": "fisharegreat",
	//	"preferredFish": null,
	//	"createdAt": "2020-08-11T23:34:50.959935+07:00"
	//}
}
