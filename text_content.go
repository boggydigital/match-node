package match_node

import (
	"golang.org/x/net/html"
	"strings"
)

func TextContent(node *html.Node) string {
	sb := &strings.Builder{}

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)

	return sb.String()
}
