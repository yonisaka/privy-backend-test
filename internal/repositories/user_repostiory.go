package repositories

import (
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type userRepo struct {
	*BaseRepo
}

func NewUserRepo(base *BaseRepo) repository.UserRepository {
	return &userRepo{BaseRepo: base}
}

func (r *userRepo) Login(ctx *gin.Context, username string) (*domain.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE username = ?"
	row, err := r.db.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}

	user := domain.User{}
	if row.Next() {
		err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
