package repository

import (
	"privy-backend-test/internal/domain"

	"github.com/gin-gonic/gin"
)

//go:generate rm -f ./user_mock.go
//go:generate moq -out ./user_mock.go . UserRepository:UserRepositoryMock
type UserRepository interface {
	Login(ctx *gin.Context, username string) (*domain.User, error)
}
