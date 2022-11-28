package usecases

import (
	"context"
	"errors"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/helpers"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type cakeUsecase struct {
	cakeRepo repository.CakeRepository
}

func NewCakeUsecase(cakeRepo repository.CakeRepository) usecase.CakeUsecase {
	return &cakeUsecase{cakeRepo: cakeRepo}
}

func (i *cakeUsecase) GetCakes(ctx context.Context) (*[]domain.Cake, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cakes, err := i.cakeRepo.GetCakes(ctx)
	if err != nil {
		return nil, err
	}

	return cakes, nil
}

func (i *cakeUsecase) GetCakeByID(ctx context.Context, id int64) (*domain.Cake, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cake, err := i.cakeRepo.GetCakeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cake, nil
}

func (i *cakeUsecase) Store(ctx context.Context, cake *domain.Cake) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := i.cakeRepo.Store(ctx, cake)
	if err != nil {
		return err
	}

	return nil
}

func (i *cakeUsecase) Update(ctx context.Context, cake *domain.Cake) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := i.cakeRepo.Update(ctx, cake)
	if err != nil {
		return err
	}

	return nil
}

func (i *cakeUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := i.cakeRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (i *cakeUsecase) UploadImage(ctx *gin.Context) (interface{}, error) {
	fileHeader, err := ctx.FormFile("file")

	if err != nil {
		return "", err
	}

	if fileHeader == nil {
		return "", errors.New("new path not found")
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	str1 := strings.Split(uuid.String(), "-")
	fn := strings.Split(fileHeader.Filename, ".")
	uniqFilename := str1[len(str1)-1:][0] + "." + fn[len(fn)-1]
	dirUpload := helpers.GetFilePath("cake")
	path := dirUpload + uniqFilename

	if err := i.cakeRepo.UploadImage(ctx, path); err != nil {
		return "", err
	}

	return map[string]interface{}{
		"filename": uniqFilename,
	}, nil
}
