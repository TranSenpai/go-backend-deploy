package storage

import "go-backend/internal/dto"

type FileStorage interface {
	Upload(input dto.UploadInput) (*dto.UploadReturn, error)
	Delete(fileName string) error
}
