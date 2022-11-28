package di

import (
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/usecases"
)

func GetCakeUsecase() usecase.CakeUsecase {
	return usecases.NewCakeUsecase(
		GetCakeRepo(),
	)
}
