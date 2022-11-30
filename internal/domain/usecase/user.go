package usecase

import (
	"privy-backend-test/internal/domain"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Login(ctx *gin.Context, user *domain.User) (*domain.Auth, error)
}
