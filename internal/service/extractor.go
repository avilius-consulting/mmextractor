// internal/service/extractor.go
package service

import (
	"errors"
	"image"
	// Register image decoders so Go can recognize JPG, PNG, and GIF headers
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"MME/internal/domain"
)

type ExtractorService struct{}

func NewExtractorService() *ExtractorService {
	return &ExtractorService{}
}

// ExtractMetadata downloads the image headers and returns dimensions and format
func (s *ExtractorService) ExtractMetadata(url string) (*domain.ImageMetadata, error) {
	// 1. Fetch the image data from the web
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch image: non-200 status code")
	}

	// 2. DecodeConfig only parses the header (super fast, low memory)
	config, format, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return nil, err
	}

	// 3. Construct and return our domain object
	return &domain.ImageMetadata{
		Format:    format,
		Width:     config.Width,
		Height:    config.Height,
		SizeBytes: resp.ContentLength, // Will show -1 if server doesn't provide Content-Length header
	}, nil
}