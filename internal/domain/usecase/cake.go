package usecase

import (
	"context"
	"privy-backend-test/internal/domain"

	"github.com/gin-gonic/gin"
)

type CakeUsecase interface {
	GetCakes(ctx context.Context) (*[]domain.Cake, error)
	GetCakeByID(ctx context.Context, id int64) (*domain.Cake, error)
	Store(ctx context.Context, cake *domain.Cake) error
	Update(ctx context.Context, cake *domain.Cake) error
	Delete(ctx context.Context, id int64) error
	UploadImage(ctx *gin.Context) (interface{}, error)
}
