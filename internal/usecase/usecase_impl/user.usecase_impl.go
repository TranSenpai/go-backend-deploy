package usecase_impl

import (
	"context"
	"fmt"
	"go-backend/ent"
	"go-backend/internal/common/pagination"
	"go-backend/internal/common/response"
	"go-backend/internal/common/storage"
	"go-backend/internal/dto"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
	"math"
	"mime/multipart"
	"path/filepath"
)

type userUsecase struct {
	userRepository        repository.UserRepository
	localFileStorage      storage.FileStorage
	cloudinaryFileStorage storage.FileStorage
}

func NewUserUsecase(userRepository repository.UserRepository, localFileStorage storage.FileStorage, cloudinaryFileStorage storage.FileStorage) usecase.UserUsecase {
	return &userUsecase{
		userRepository:        userRepository,
		localFileStorage:      localFileStorage,
		cloudinaryFileStorage: cloudinaryFileStorage,
	}
}

// FindAll implements [usecase.UserUsecase].
func (a *userUsecase) FindAll(ctx context.Context, input dto.UserFindAllInput) (any, error) {
	data, err := a.userRepository.GetAll(ctx, input.Query, input.UserFindAllFilters)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// totalItem: tổng số lượng item
	totalItem, err := a.userRepository.Count(ctx, input.UserFindAllFilters)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	// totalPage: tổng số trang totalItem / pageSize
	totalPage := float64(totalItem) / float64(input.PageSize)

	result := pagination.PaginationRes[any]{
		Items:     data,
		Page:      input.Page,
		PageSize:  input.PageSize,
		TotalItem: totalItem,
		TotalPage: int(math.Ceil(totalPage)),
	}

	return result, nil
}

// FindOne implements [usecase.UserUsecase].
func (a *userUsecase) FindOne(ctx context.Context, id int) (any, error) {
	return a.userRepository.FindUserById(ctx, id)
}

// AvatarLocal implements [usecase.UserUsecase].
func (a *userUsecase) AvatarLocal(ctx context.Context, fileHeader *multipart.FileHeader, user *ent.Users) (any, error) {
	input := dto.UploadInput{
		FileHeader: fileHeader,
		Folder:     "images",
		Prefix:     "local-",
	}
	resultUpload, err := a.localFileStorage.Upload(input)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}

	fmt.Printf("\n\n %+v \n\n", resultUpload)

	userUpdate, err := a.userRepository.UpdateAvatarById(ctx, user.ID, resultUpload.FileName)
	if err != nil {
		// xoá tấm hình nếu lưu db lỗi
		err := a.localFileStorage.Delete(filepath.Join(input.Folder, resultUpload.FileName))
		if err != nil {
			fmt.Println("localFileStorage.Delete", err.Error())
		}
		err = a.cloudinaryFileStorage.Delete(resultUpload.FullPath)
		if err != nil {
			fmt.Println("cloudinaryFileStorage.Delete", err.Error())
		}
		return nil, response.NewBadRequestException(err.Error())
	}

	if user.Avatar != nil {
		// xoá hình cũ

		// filepath.Join() // file
		// path.Join() // url
		err := a.localFileStorage.Delete(filepath.Join(input.Folder, *user.Avatar))
		if err != nil {
			fmt.Println("localFileStorage.Delete", err.Error())
		}
		err = a.cloudinaryFileStorage.Delete(*user.Avatar)
		if err != nil {
			fmt.Println("cloudinaryFileStorage.Delete", err.Error())
		}
	}

	return userUpdate, nil
}

// AvatarCloud implements [usecase.UserUsecase].
func (a *userUsecase) AvatarCloud(ctx context.Context, fileHeader *multipart.FileHeader, user *ent.Users) (any, error) {
	input := dto.UploadInput{
		FileHeader: fileHeader,
		Folder:     "images",
		Prefix:     "cloud-",
	}
	resultUpload, err := a.cloudinaryFileStorage.Upload(input)
	if err != nil {
		return nil, response.NewBadRequestException(err.Error())
	}
	fmt.Println("resultUpload", resultUpload)

	userUpdate, err := a.userRepository.UpdateAvatarById(ctx, user.ID, resultUpload.FullPath)
	if err != nil {
		// xoá tấm hình nếu lưu db lỗi
		err := a.cloudinaryFileStorage.Delete(resultUpload.FullPath)
		if err != nil {
			fmt.Println("cloudinaryFileStorage.Delete", err.Error())
		}
		err = a.localFileStorage.Delete(filepath.Join(input.Folder, resultUpload.FileName))
		if err != nil {
			fmt.Println("localFileStorage.Delete", err.Error())
		}
		return nil, response.NewBadRequestException(err.Error())
	}

	if user.Avatar != nil {
		// xoá hình cũ

		// filepath.Join() // file
		// path.Join() // url
		err := a.cloudinaryFileStorage.Delete(*user.Avatar)
		if err != nil {
			fmt.Println("cloudinaryFileStorage.Delete", err.Error())
		}
		err = a.localFileStorage.Delete(filepath.Join(input.Folder, *user.Avatar))
		if err != nil {
			fmt.Println("localFileStorage.Delete", err.Error())
		}
	}

	return userUpdate, nil
}
