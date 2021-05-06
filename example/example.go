package main

import (
	"fmt"
	"os"

	goseo "github.com/dant89/go-seo"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a target URL")
		os.Exit(1)
	}

	url := os.Args[1]
	webpage := goseo.Webpage{Url: url}

	response, err := webpage.Crawl()
	if err != nil {
		fmt.Println("Could not crawl the URL:", err.Error())
		os.Exit(2)
	}

	h1, err := goseo.GetFirstElement(response, "h1")
	if err != nil {
		fmt.Println("Could not parse a h1 in the HTML:", err.Error())
		os.Exit(3)
	}

	h1LengthStatus := goseo.CheckH1Length(h1)

	fmt.Println("SEO advice for:", url)
	fmt.Printf("H1 currently: '%s'\n", h1)
	fmt.Println("H1 status:", h1LengthStatus)
}
