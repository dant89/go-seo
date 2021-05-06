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

	h2Count := goseo.GetAllElements(response, "h2")
	fmt.Println("H2 count:", len(h2Count))

	h3Count := goseo.GetAllElements(response, "h3")
	fmt.Println("H3 count:", len(h3Count))

	h4Count := goseo.GetAllElements(response, "h4")
	fmt.Println("H4 count:", len(h4Count))

	aCount := goseo.GetAllElements(response, "a")
	fmt.Println("Link count:", len(aCount))
}
