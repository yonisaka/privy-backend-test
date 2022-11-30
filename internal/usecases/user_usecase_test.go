package usecases_test

import (
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/usecases"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type userFields struct {
	userRepo repository.UserRepository
}

func userSut(f userFields) usecase.UserUsecase {
	return usecases.NewUserUsecase(f.userRepo)
}

func TestLogin(t *testing.T) {
	type args struct {
		ctx  *gin.Context
		user *domain.User
	}

	type test struct {
		fields  userFields
		args    args
		want    *domain.Auth
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling user repo Login succeed, Should return no error": func(t *testing.T) test {
			ctx := gin.Context{}

			user := &domain.User{
				Username: "userdemo",
				Password: "password",
			}

			args := args{
				ctx:  &ctx,
				user: user,
			}

			return test{
				fields: userFields{
					userRepo: &repository.UserRepositoryMock{
						LoginFunc: func(ctx *gin.Context, username string) (*domain.User, error) {
							assert.Equal(t, args.ctx, ctx)

							return &domain.User{
								ID:       1,
								Username: "userdemo",
								Email:    "user@gmail.com",
								Password: "$2a$14$WjHP5kc8tE0LQzVuTqlwNehjNbYpUgz1f7hgv.98XWd8mgoJMMzFa",
							}, nil
						},
					},
				},
				args: args,
				want: &domain.Auth{
					Access_token: "access_token",
					Expired_at:   "expired_at",
				},
				wantErr: nil,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := userSut(tt.fields)

			_, err := sut.Login(tt.args.ctx, tt.args.user)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
