package extractlinks

import (
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

// Link object for parsing an anchor link
type Link struct {
	Href string
	Text string
}

// All takes a reader object (like the one returned from http.Get())
// It returns a slice of Links representing the Href & Text attributes from
// anchor links found in the provided html.
// It does not close the reader passed to it.
func All(htmlBody io.Reader) ([]Link, error) {
	document, err := html.Parse(htmlBody)
	if err != nil {
		return nil, err
	}

	nodes := buildNodes(document)

	var links []Link
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}

	return links, nil
}

func buildNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, buildNodes(c)...)
	}
	return ret
}

func buildLink(n *html.Node) (link Link) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = trimHash(attr.Val)
		}
	}

	link.Text = buildText(n)
	return
}

func buildText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += buildText(c)
	}
	return strings.Join(strings.Fields(text), " ")
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}
