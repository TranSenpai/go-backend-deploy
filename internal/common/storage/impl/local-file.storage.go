package storage_impl

import (
	"fmt"
	"go-backend/internal/common/storage"
	"go-backend/internal/dto"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type localFileStorage struct {
	baseFolder string
}

func NewLocalFileStorage(baseFolder string) storage.FileStorage {
	return &localFileStorage{
		baseFolder: baseFolder,
	}
}

// Upload implements [storage.FileStorage].
func (l *localFileStorage) Upload(input dto.UploadInput) (*dto.UploadReturn, error) {
	ext := strings.ToLower(filepath.Ext(input.FileHeader.Filename))
	fmt.Println("ext", ext)

	fileName := fmt.Sprintf("%s%d%s", input.Prefix, time.Now().UnixNano(), ext)
	fmt.Println("fileName", fileName)

	fullPath := filepath.Join(l.baseFolder, input.Folder, fileName)
	fmt.Println("fullPath", fullPath)

	err := saveUploadedFile(input.FileHeader, fullPath)
	if err != nil {
		return nil, err
	}

	return &dto.UploadReturn{
		FileName: fileName,
		FullPath: fullPath,
		Url:      "",
	}, nil
}

func saveUploadedFile(file *multipart.FileHeader, dst string, perm ...fs.FileMode) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var mode os.FileMode = 0o750
	if len(perm) > 0 {
		mode = perm[0]
	}
	dir := filepath.Dir(dst)
	if err = os.MkdirAll(dir, mode); err != nil {
		return err
	}
	if err = os.Chmod(dir, mode); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// Delete implements [storage.FileStorage].
func (l *localFileStorage) Delete(fileName string) error {
	err := os.Remove(filepath.Join(l.baseFolder, fileName))

	if err != nil {
		return err
	}

	return nil
}
