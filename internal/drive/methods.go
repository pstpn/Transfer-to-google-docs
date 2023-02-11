package drive

import (
	"context"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/drive/v3"
)

// Service is a Google Drive service.
type Service struct {
	client  *http.Client   // client is the HTTP client used to communicate with the API.
	service *drive.Service // service is the Drive service.
}

// NewService creates a new Google Drive service.
func NewService(ctx context.Context, secretFile string) (*Service, error) {

	b, err := os.ReadFile(secretFile)
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
		return nil, err
	}

	client := getClient(ctx, config)

	service, err := drive.New(client)
	if err != nil {
		log.Printf("Unable to retrieve Drive client: %v", err)
		return nil, err
	}

	return &Service{
		client:  client,
		service: service,
	}, nil
}

// CreateFileInDrive creates a file in Google Drive.
func (s *Service) CreateFileInDrive(dataFilename string) (string, error) {

	f, err := os.Open(dataFilename)
	if err != nil {
		log.Printf("Error opening %q: %v", dataFilename, err)
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Error closing %q: %v", dataFilename, err)
		}
	}()

	driveFile, err := s.service.Files.Create(&drive.File{Name: dataFilename}).Media(f).Do()
	if err != nil {
		log.Printf("Unable to create file: %v", err)
		return "", err
	}

	log.Printf("File \"%s\" uploaded with ID \"%s\"", driveFile.Name, driveFile.Id)

	return driveFile.Id, nil
}
