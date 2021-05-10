package goseo

import (
	"bytes"
	"errors"
	"io"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
	"golang.org/x/net/html"
)

type Parser struct{}

type ParserChecks interface {
	GetAllLinkHrefs()
	GetFirstElement()
	GetAllElements()
}

func (p Parser) GetAllLinkHrefs(rawHtml string) []string {
	var formattedElements []string
	doc, _ := html.Parse(strings.NewReader(rawHtml))

	// recursively parse the html until we find a H1
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" && (strings.HasPrefix(a.Val, "/") || strings.HasPrefix(a.Val, "http")) {
					formattedElements = append(formattedElements, a.Val)
					break
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)

	return formattedElements
}

func (p Parser) GetFirstElement(rawHtml string, element string, raw bool) (string, error) {
	var foundElement *html.Node
	doc, _ := html.Parse(strings.NewReader(rawHtml))

	// recursively parse the html until we find a H1
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == element {
			foundElement = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if foundElement != nil {
		formatted := renderNode(foundElement)
		if !raw {
			formatted = stripHtml(formatted)
		}
		return formatted, nil
	}

	return "", errors.New("missing <" + element + "> in the node tree")
}

func (p Parser) GetAllElements(rawHtml string, element string, raw bool) []string {
	var foundElements []*html.Node
	var formattedElements []string
	doc, _ := html.Parse(strings.NewReader(rawHtml))

	// recursively parse the html until we find a H1
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == element {
			foundElements = append(foundElements, node)
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)

	if len(foundElements) > 0 {
		for _, v := range foundElements {
			formatted := renderNode(v)
			if !raw {
				formatted = stripHtml(formatted)
			}
			formattedElements = append(formattedElements, formatted)
		}
	}

	return formattedElements
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func stripHtml(h1 string) string {
	stripped := strip.StripTags(h1)
	trimmed := strings.TrimSpace(stripped)
	return trimmed
}
