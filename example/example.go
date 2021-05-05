package main

import (
	"fmt"

	goseo "github.com/dant89/go-seo"
)

func main() {
	// TODO convert URL to command line argument
	var url string = "https://dant.blog"
	webpage := goseo.Webpage{Url: url}
	response, _ := webpage.Crawl()
	h1, _ := goseo.GetH1(response)
	h1LengthStatus := goseo.CheckH1Length(h1)
	fmt.Println("SEO advice for:", url)
	fmt.Printf("H1 currently: '%s'\n", h1)
	fmt.Println("H1 status:", h1LengthStatus)
}
