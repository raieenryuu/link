package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHtml(r io.Reader) ([]Link, error) {
	node, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	return GetAllLinks(node), nil
}

func GetAllLinks(node *html.Node) []Link {
	var result []Link

	var f func(n *html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					result = append(result, Link{Href: attr.Val, Text: extractText(n)})
					break
				}

				fmt.Println(attr.Key)
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)

	return result
}

func extractText(node *html.Node) string {
	result := ""

	var f func(n *html.Node)

	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			result += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}

	f(node)

	return strings.Join(strings.Fields(result), " ")
}
