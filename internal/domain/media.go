package domain

type ImageMetadata struct {
	Format string `json:"format"`
	Width  int	  `json:"width"`
	Height int    `json:"height"`
	SizeBytes int64 `json:"size_bytes"`
}