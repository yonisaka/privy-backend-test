package usecases

import (
	"errors"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/helpers"

	"github.com/gin-gonic/gin"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (i *userUsecase) Login(ctx *gin.Context, user *domain.User) (*domain.Auth, error) {
	res, err := i.userRepo.Login(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPasswordHash(user.Password, res.Password) {
		return nil, errors.New("username or password is not match")
	}

	jwt := helpers.CreateTokenJWT(res.ID)

	auth := domain.Auth{
		Access_token: jwt.AccessToken,
		Expired_at:   helpers.ConvertToString(jwt.AtExpires),
	}

	return &auth, nil
}
