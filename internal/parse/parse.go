package parse

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

// GetHTML gets the HTML document from the URL.
func GetHTML(url string) (*http.Response, error) {

	doc, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// GetTableData gets the table data from the HTML document.
func GetTableData(body io.Reader) ([][]string, error) {

	table := make([][]string, 0)

	// Parse the HTML document.
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	// Find the table tag.
	tableNode, err := findTableTag(doc)
	if err != nil {
		return nil, err
	}

	// Get the table title.
	tableTitle, err := getTableTitle(tableNode)
	if err != nil {
		return nil, err
	}

	table = append(table, tableTitle)

	// Get the table.
	for c := tableNode.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "tbody" {
			for r := c.FirstChild; r != nil; r = r.NextSibling {
				row := make([]string, 0)

				for d := r.FirstChild; d != nil; d = d.NextSibling {
					if d.Data == "td" {
						rowPart, _ := getRow(d)
						row = append(row, rowPart...)
					}
				}

				table = append(table, row)
			}
		}
	}

	return table, nil
}

// WriteTableToFile writes the table to a file.
func WriteTableToFile(filename string, table [][]string) error {

	// Create a new file.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()

	// Write the table to the file.
	for _, row := range table {
		for i := 0; i < len(table[0]); i++ {
			if i == 1 {
				for j := 1; j < len(row); j++ {
					_, err := file.WriteString(row[j] + " ")
					if err != nil {
						return err
					}
				}

				_, err := file.WriteString(";")
				if err != nil {
					return err
				}
			} else {
				_, err := file.WriteString(row[i] + ";")
				if err != nil {
					return err
				}
			}
		}

		_, err := file.WriteString("\n")
		if err != nil {
			return err
		}
	}

	return nil
}
