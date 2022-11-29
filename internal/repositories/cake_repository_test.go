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

// var ctx *gin.Context

func TestGetCakes(t *testing.T) {
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
	mockCakes := []domain.Cake{
		{
			ID:          1,
			Title:       "Lemon Cheesecake",
			Description: "A cheesecake made of lemon",
			Rating:      7,
			Image:       "no-image.jpeg",
			CreatedAt:   &timeNow,
			UpdatedAt:   &timeNow,
		},
		{
			ID:          2,
			Title:       "Cake cake",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam",
			Rating:      5.5,
			Image:       "no-image.jpeg",
			CreatedAt:   &timeNow,
			UpdatedAt:   &timeNow,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(mockCakes[0].ID, mockCakes[0].Title, mockCakes[0].Description, mockCakes[0].Rating, mockCakes[0].Image, mockCakes[0].CreatedAt, mockCakes[0].UpdatedAt).
		AddRow(mockCakes[1].ID, mockCakes[1].Title, mockCakes[1].Description, mockCakes[1].Rating, mockCakes[1].Image, mockCakes[1].CreatedAt, mockCakes[1].UpdatedAt)

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := di.GetCakeRepo()

	_, err = repo.GetCakes(ctx)
	assert.NoError(t, err)
}

func TestGetCakeByID(t *testing.T) {
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
	mockCakes := []domain.Cake{
		{
			ID:          1,
			Title:       "Lemon Cheesecake",
			Description: "A cheesecake made of lemon",
			Rating:      7,
			Image:       "no-image.jpeg",
			CreatedAt:   &timeNow,
			UpdatedAt:   &timeNow,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(mockCakes[0].ID, mockCakes[0].Title, mockCakes[0].Description, mockCakes[0].Rating, mockCakes[0].Image, mockCakes[0].CreatedAt, mockCakes[0].UpdatedAt)

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := di.GetCakeRepo()
	id := int64(1)
	cake, err := repo.GetCakeByID(ctx, id)
	assert.NoError(t, err)
	assert.NotNil(t, cake)
	assert.Equal(t, mockCakes[0].ID, cake.ID)
}

func TestCreateCake(t *testing.T) {
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
	mockCake := domain.Cake{
		Title:       "Lemon Cheesecake",
		Description: "A cheesecake made of lemon",
		Rating:      7,
		Image:       "no-image.jpeg",
		CreatedAt:   &timeNow,
		UpdatedAt:   &timeNow,
	}

	query := "INSERT INTO cakes \\(title, description, rating, image, created_at, updated_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockCake.Title, mockCake.Description, mockCake.Rating, mockCake.Image, mockCake.CreatedAt, mockCake.UpdatedAt).WillReturnResult(sqlmock.NewResult(3, 1))
	repo := di.GetCakeRepo()
	err = repo.Store(ctx, &mockCake)
	assert.NoError(t, err)
}

func TestUpdateCake(t *testing.T) {
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
	mockCake := domain.Cake{
		ID:          1,
		Title:       "Lemon Update",
		Description: "A cheesecake made of lemon",
		Rating:      7,
		Image:       "no-image.jpeg",
		CreatedAt:   &timeNow,
		UpdatedAt:   &timeNow,
	}

	query := "UPDATE cakes SET title = \\?, description = \\?, rating = \\?, image = \\?, updated_at = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockCake.Title, mockCake.Description, mockCake.Rating, mockCake.Image, mockCake.UpdatedAt, mockCake.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := di.GetCakeRepo()
	err = repo.Update(ctx, &mockCake)
	assert.NoError(t, err)
}

func TestDeleteCake(t *testing.T) {
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

	query := "DELETE FROM cakes WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := di.GetCakeRepo()
	err = repo.Delete(ctx, 2)
	assert.NoError(t, err)
}
