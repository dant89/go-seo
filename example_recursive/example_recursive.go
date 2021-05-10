package main

import (
	"fmt"
	"os"
	"time"

	goseo "github.com/dant89/go-seo"
)

var spider goseo.Spider = goseo.Spider{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a target URL")
		os.Exit(1)
	}

	url := os.Args[1]
	getInternalLinks(url)
	getAllLinks(url)
}

func getInternalLinks(url string) {
	start := time.Now()
	internalLinks := spider.GetLinks(url, true, 2)
	if len(internalLinks) > 0 {
		fmt.Println("Internal links found:")
		fmt.Println(">", len(internalLinks), "links")
		fmt.Println("> with a search depth of 2")
		fmt.Printf("> taking %v\n", time.Since(start))
	} else {
		fmt.Println("No internal links found.")
	}
}

func getAllLinks(url string) {
	start := time.Now()
	allLinks := spider.GetLinks(url, false, 2)
	if len(allLinks) > 0 {
		fmt.Println("Internal and external links found:")
		fmt.Println(">", len(allLinks), "links")
		fmt.Println("> with a search depth of 2")
		fmt.Printf("> taking %v\n", time.Since(start))
	} else {
		fmt.Println("No internal or external links found.")
	}
}
