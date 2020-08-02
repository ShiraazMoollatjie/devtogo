package main

import (
	"fmt"
	"log"

	"github.com/ShiraazMoollatjie/devtogo"
)

func main() {
	cl := devtogo.NewClient(devtogo.WithApiKey("YOUR_API_KEY_HERE"))
	al, err := cl.Articles(devtogo.Defaults())
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("All Articles: %+v", al)

	a, err := cl.PublishedArticle(416009)
	if err != nil {
		log.Fatalf("something went wrong: %+v", err)
	}

	fmt.Printf("Article: %+v", a.ID)
}
