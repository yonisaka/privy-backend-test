package repositories

import (
	"context"
	"fmt"
	"os"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type cakeRepo struct {
	*BaseRepo
}

func NewCakeRepo(base *BaseRepo) repository.CakeRepository {
	return &cakeRepo{BaseRepo: base}
}

func (i *cakeRepo) GetCakes(ctx context.Context) (*[]domain.Cake, error) {
	var cake []domain.Cake
	err := i.db.WithContext(ctx).Find(&cake).Error
	if err != nil {
		return nil, err
	}
	return &cake, nil
}

func (i *cakeRepo) GetCakeByID(ctx context.Context, id int64) (*domain.Cake, error) {
	var cake domain.Cake
	err := i.db.WithContext(ctx).First(&cake, id).Error
	if err != nil {
		return nil, err
	}
	return &cake, nil
}

func (i *cakeRepo) Store(ctx context.Context, cake *domain.Cake) error {
	err := i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Create(&cake).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (i *cakeRepo) Update(ctx context.Context, cake *domain.Cake) error {
	err := i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		row := tx.WithContext(ctx).Where("id", cake.ID).First(&domain.Cake{})
		if row.RowsAffected == 0 {
			return fmt.Errorf("cake not found")
		}

		err := tx.WithContext(ctx).Updates(&cake).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (i *cakeRepo) Delete(ctx context.Context, id int64) error {
	err := i.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		cake := domain.Cake{}
		row := tx.WithContext(ctx).Where("id", id).First(&cake)
		if row.RowsAffected == 0 {
			return fmt.Errorf("cake not found")
		}
		if err := tx.WithContext(ctx).Where("id", id).Delete(&cake).Error; err != nil {
			return err
		}
		if cake.Image != "" {
			filePath := helpers.GetFilePath("cake")
			if err := os.Remove(filePath + cake.Image); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (i *cakeRepo) UploadImage(ctx *gin.Context, path string) error {
	file := ctx.Request.MultipartForm.File["file"][0]

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		return err
	}
	return nil
}
