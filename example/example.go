package main

import (
	"fmt"
	"os"

	goseo "github.com/dant89/go-seo"
)

var spider goseo.Spider = goseo.Spider{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a target URL")
		os.Exit(1)
	}

	url := os.Args[1]
	analyser := goseo.Analyse{}
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

	h1Report := analyser.CheckH1Length(h1)
	fmt.Println("SEO advice for:", url)
	fmt.Printf("\nH1 currently: '%s'\n", h1)
	if !h1Report.Passed() {
		for _, report := range h1Report.GetFeedback() {
			fmt.Println("H1 feedback:", report)
		}
	}

	h2Count := parser.GetAllElements(response, "h2", false)
	fmt.Println("H2 count:", len(h2Count))

	h3Count := parser.GetAllElements(response, "h3", false)
	fmt.Println("H3 count:", len(h3Count))

	h4Count := parser.GetAllElements(response, "h4", false)
	fmt.Println("H4 count:", len(h4Count))

	getLinks(url)
}

func getLinks(url string) {
	links := spider.GetLinks(url, 1)

	internalLinkCount := 0
	for _, link := range links {
		if link.Internal {
			internalLinkCount++
		}
	}

	externalLinkCount := 0
	for _, link := range links {
		if !link.Internal {
			externalLinkCount++
		}
	}

	fmt.Println("Internal links found:", internalLinkCount)
	fmt.Println("External links found:", externalLinkCount)
}
