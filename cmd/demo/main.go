package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cesbo/helpscout-toolkit/pkg/helpscout"
)

var (
	apiKey = os.Getenv("HELPSCOUT_API_KEY")
)

func main() {
	var (
		err  error
		data []byte
	)

	hs := helpscout.NewHelpScout(apiKey)
	summary, err := hs.GetSummary()
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.MarshalIndent(summary, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
