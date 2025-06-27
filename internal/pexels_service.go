package internal

import (
	"context"
	"note-taker/pexels"
)

type PexelsService struct{}

func NewPexelService() *PexelsService {
	return &PexelsService{}
}

func (s *PexelsService) GetPictures(ctx context.Context, query string) ([]pexels.Photo, error) {
	photos, err := pexels.GetPhotos(query)
	return photos, err
}
