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
	webpage := goseo.Crawler{}
	parser := goseo.Parser{}

	response, err := webpage.Crawl(url)
	if err != nil {
		fmt.Println("Could not crawl the URL:", err.Error())
		os.Exit(2)
	}

	h1, err := parser.GetFirstElement(response, "h1", false)
	if err != nil {
		fmt.Println("Could not parse a h1 in the HTML:", err.Error())
		os.Exit(3)
	}

	// TODO tidy globally used function
	h1LengthStatus := goseo.CheckH1Length(h1)

	fmt.Println("SEO advice for:", url)
	fmt.Printf("H1 currently: '%s'\n", h1)
	fmt.Println("H1 status:", h1LengthStatus)

	h2Count := parser.GetAllElements(response, "h2", false)
	fmt.Println("H2 count:", len(h2Count))

	h3Count := parser.GetAllElements(response, "h3", false)
	fmt.Println("H3 count:", len(h3Count))

	h4Count := parser.GetAllElements(response, "h4", false)
	fmt.Println("H4 count:", len(h4Count))

	aCount := parser.GetAllElements(response, "a", false)
	fmt.Println("Link count:", len(aCount))
}
