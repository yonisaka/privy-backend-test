package repository

import (
	"privy-backend-test/internal/domain"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Login(ctx *gin.Context, username string) (*domain.User, error)
}
