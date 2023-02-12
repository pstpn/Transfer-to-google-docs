package parse

import (
	"errors"
	"golang.org/x/net/html"
)

// findTableTag finds the first table tag in the HTML document.
func findTableTag(n *html.Node) (*html.Node, error) {

	// Check if the tag is the one we are looking for.
	if n.Type == html.ElementNode && n.Data == "table" {
		return n, nil
	}

	// Recursively search for the tag.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		curN, _ := findTableTag(c)

		if curN != nil {
			return curN, nil
		}
	}

	return nil, errors.New("table tag not found")
}

// getRow gets the row from the HTML table.
func getRow(node *html.Node) ([]string, error) {

	var row []string

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			row = append(row, c.Data)
		} else if c.FirstChild != nil {
			curRow, _ := getRow(c)
			row = append(row, curRow...)
		}
	}

	return row, nil
}

// getTableTitle gets the table title from the HTML document.
func getTableTitle(tableNode *html.Node) ([]string, error) {

	title := make([]string, 0)

	// Get the table title.
	for c := tableNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "thead" {
			c = c.FirstChild.FirstChild

			for ; c != nil; c = c.NextSibling {
				if c.Data == "th" {
					title = append(title, c.FirstChild.Data)
				}
			}

			break
		}
	}

	return title, nil
}
