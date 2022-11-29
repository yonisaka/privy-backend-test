package repository

import (
	"privy-backend-test/internal/domain"

	"github.com/gin-gonic/gin"
)

//go:generate rm -f ./cake_mock.go
//go:generate moq -out ./cake_mock.go . CakeRepository:CakeRepositoryMock
type CakeRepository interface {
	GetCakes(ctx *gin.Context) (*[]domain.Cake, error)
	GetCakeByID(ctx *gin.Context, id int64) (*domain.Cake, error)
	Store(ctx *gin.Context, cake *domain.Cake) error
	Update(ctx *gin.Context, cake *domain.Cake) error
	Delete(ctx *gin.Context, id int64) error
	UploadImage(ctx *gin.Context, path string) error
}
