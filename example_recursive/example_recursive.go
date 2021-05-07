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
	spider := goseo.Spider{}

	internalLinks := spider.GetLinks(url, true)
	if len(internalLinks) > 0 {
		fmt.Println("Internal links found:")
		for _, link := range internalLinks {
			fmt.Println(link)
		}
	} else {
		fmt.Println("No internal links found.")
	}
}
