// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"github.com/gin-gonic/gin"
	"privy-backend-test/internal/domain"
	"sync"
)

// Ensure, that UserRepositoryMock does implement UserRepository.
// If this is not the case, regenerate this file with moq.
var _ UserRepository = &UserRepositoryMock{}

// UserRepositoryMock is a mock implementation of UserRepository.
//
//	func TestSomethingThatUsesUserRepository(t *testing.T) {
//
//		// make and configure a mocked UserRepository
//		mockedUserRepository := &UserRepositoryMock{
//			LoginFunc: func(ctx *gin.Context, username string) (*domain.User, error) {
//				panic("mock out the Login method")
//			},
//		}
//
//		// use mockedUserRepository in code that requires UserRepository
//		// and then make assertions.
//
//	}
type UserRepositoryMock struct {
	// LoginFunc mocks the Login method.
	LoginFunc func(ctx *gin.Context, username string) (*domain.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// Login holds details about calls to the Login method.
		Login []struct {
			// Ctx is the ctx argument value.
			Ctx *gin.Context
			// Username is the username argument value.
			Username string
		}
	}
	lockLogin sync.RWMutex
}

// Login calls LoginFunc.
func (mock *UserRepositoryMock) Login(ctx *gin.Context, username string) (*domain.User, error) {
	if mock.LoginFunc == nil {
		panic("UserRepositoryMock.LoginFunc: method is nil but UserRepository.Login was just called")
	}
	callInfo := struct {
		Ctx      *gin.Context
		Username string
	}{
		Ctx:      ctx,
		Username: username,
	}
	mock.lockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	mock.lockLogin.Unlock()
	return mock.LoginFunc(ctx, username)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//
//	len(mockedUserRepository.LoginCalls())
func (mock *UserRepositoryMock) LoginCalls() []struct {
	Ctx      *gin.Context
	Username string
} {
	var calls []struct {
		Ctx      *gin.Context
		Username string
	}
	mock.lockLogin.RLock()
	calls = mock.calls.Login
	mock.lockLogin.RUnlock()
	return calls
}