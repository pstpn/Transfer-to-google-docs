package main

import (
	"context"
	"log"
	"parse/config"
	"parse/internal/drive"
	"parse/internal/parse"
)

// main is the key function of the project.
func main() {

	// Get HTML from the page.
	pageHTML, err := parse.GetHTML(config.ParseURL)
	if err != nil {
		log.Printf("Unable to get page \"%s\" : %v", config.ParseURL, err)
		return
	}
	defer func() {
		err := pageHTML.Body.Close()
		if err != nil {
			log.Printf("Unable to close page \"%s\" : %v", config.ParseURL, err)
		}
	}()

	log.Printf("Page parsed successfully")

	// Parse HTML.
	table, err := parse.GetTableData(pageHTML.Body)
	if err != nil {
		log.Printf("Unable to parse page \"%s\" : %v", config.ParseURL, err)
		return
	}

	log.Printf("Table parsed successfully")

	// Write the table to a file.
	err = parse.WriteTableToFile(config.OutputFilename, table)

	// Create a new Drive service.
	service, err := drive.NewService(context.Background(), config.CredentialsFile)
	if err != nil {
		log.Printf("Unable to retrieve Drive client: %v", err)
		return
	}

	log.Printf("Drive service created successfully")

	// Create a file in the Drive.
	_, err = service.CreateFileInDrive(config.OutputFilename)
	if err != nil {
		log.Printf("Unable to create file: %v", err)
		return
	}

	log.Printf("File created successfully")

}
