package main

import (
	"context"
	"log"
	"parse/config"
	"parse/internal/drive"
)

// main is the key function of the project.
func main() {

	service, err := drive.NewService(context.Background(), config.CredentialsFile)
	if err != nil {
		log.Printf("Unable to retrieve Drive client: %v", err)
		return
	}

	//_, err = service.CreateFileInDrive("data/data.txt")
	//if err != nil {
	//	log.Printf("Unable to create file: %v", err)
	//	return
	//}
}
