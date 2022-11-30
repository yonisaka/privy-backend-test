package repositories_test

import (
	"net/http/httptest"
	"privy-backend-test/internal/di"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/helpers"
	"privy-backend-test/internal/repositories"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	repositories.NewBaseRepo(sqlxDB)
	defer mockDB.Close()

	timeNow, err := helpers.ConvertToTime(time.Now())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when convert time", err)
	}
	mockUsers := []domain.User{
		{
			ID:        1,
			Username:  "userdemo",
			Email:     "user@gmail.com",
			Password:  "$2a$14$WjHP5kc8tE0LQzVuTqlwNehjNbYpUgz1f7hgv.98XWd8mgoJMMzFa",
			CreatedAt: &timeNow,
			UpdatedAt: &timeNow,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "created_at", "updated_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Username, mockUsers[0].Email, mockUsers[0].Password, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt)

	query := "SELECT id, username, email, password FROM users WHERE username = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := di.GetUserRepo()
	username := mockUsers[0].Username
	cake, err := repo.Login(ctx, username)
	assert.NoError(t, err)
	assert.NotNil(t, cake)
	assert.Equal(t, mockUsers[0].ID, cake.ID)
}
