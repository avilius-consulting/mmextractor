// internal/domain/media.go
package domain

// ImageMetadata represents the extracted characteristics of an image file.
type ImageMetadata struct {
	Format string `json:"format"`
	Width  int	  `json:"width"`
	Height int    `json:"height"`
	SizeBytes int64 `json:"size_bytes"`
}

// ExtractRequest defines the expected JSON payload from upstream clients.
type ExtractRequest struct {
	ImageURL string `json:"image_url"`
}