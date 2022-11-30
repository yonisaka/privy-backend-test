package di

import (
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/mysql"
	"privy-backend-test/internal/repositories"
)

func GetCakeRepo() repository.CakeRepository {
	return repositories.NewCakeRepo(GetBaseRepo())
}

func GetUserRepo() repository.UserRepository {
	return repositories.NewUserRepo(GetBaseRepo())
}

func GetBaseRepo() *repositories.BaseRepo {
	return repositories.NewBaseRepo(mysql.Connection())
}
