package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cesbo/helpscout-toolkit/pkg/helpscout"
)

func main() {
	var (
		err  error
		data []byte
	)

	fmt.Println("List collections")

	var collections *helpscout.Collections
	collections, err = helpscout.ListCollections()
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.Marshal(collections)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	if len(collections.Items) == 0 {
		return
	}

	fmt.Println("List categories in", collections.Items[0].Name)

	// get categories from first collection
	var categories *helpscout.Categories
	categories, err = helpscout.ListCategories(collections.Items[0].Id)
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.Marshal(categories)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println("List articles in", categories.Items[0].Name)

	// get articles from first category
	var articles *helpscout.Articles
	articles, err = helpscout.ListArticles(categories.Items[0].Id)
	if err != nil {
		log.Fatal(err)
	}

	data, err = json.Marshal(articles)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
