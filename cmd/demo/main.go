package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cesbo/helpscout-toolkit/pkg/helpscout"
)

func main() {
	collections, err := helpscout.ListCollections()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(collections)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
