package di

import "privy-backend-test/internal/adapters/handler"

func GetCakeHandler() *handler.CakeHandler {
	return handler.NewCakeHandler(GetCakeUsecase())
}

func GetUserHandler() *handler.UserHandler {
	return handler.NewUserHandler(GetUserUsecase())
}
