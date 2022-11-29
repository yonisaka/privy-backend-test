package repositories

import (
	"errors"
	"os"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/helpers"

	"github.com/gin-gonic/gin"
)

type cakeRepo struct {
	*BaseRepo
}

func NewCakeRepo(base *BaseRepo) repository.CakeRepository {
	return &cakeRepo{BaseRepo: base}
}

func (r *cakeRepo) GetCakes(ctx *gin.Context) (*[]domain.Cake, error) {
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes"
	res, err := r.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *cakeRepo) GetCakeByID(ctx *gin.Context, id int64) (*domain.Cake, error) {
	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"
	res, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(*res) == 0 {
		return nil, errors.New("cake not found")
	}

	return &(*res)[0], nil
}

func (r *cakeRepo) fetch(ctx *gin.Context, query string, args ...interface{}) (*[]domain.Cake, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cakes []domain.Cake
	for rows.Next() {
		cake := domain.Cake{}
		err := rows.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return &cakes, nil
}

func (r *cakeRepo) Store(ctx *gin.Context, cake *domain.Cake) error {
	store, err := r.db.PrepareContext(ctx, "INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer store.Close()

	_, err = store.ExecContext(ctx, cake.Title, cake.Description, cake.Rating, cake.Image, cake.CreatedAt, cake.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *cakeRepo) Update(ctx *gin.Context, cake *domain.Cake) error {
	cakeExist, err := r.GetCakeByID(ctx, cake.ID)
	if err != nil {
		return err
	}

	if cakeExist == nil {
		return errors.New("cake not found")
	}

	update, err := r.db.PrepareContext(ctx, "UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer update.Close()

	_, err = update.ExecContext(ctx, cake.Title, cake.Description, cake.Rating, cake.Image, cake.UpdatedAt, cake.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *cakeRepo) Delete(ctx *gin.Context, id int64) error {
	cakeExist, err := r.GetCakeByID(ctx, id)
	if err != nil {
		return err
	}

	if cakeExist == nil {
		return errors.New("cake not found")
	}

	if cakeExist.Image != "" && cakeExist.Image != "no-image.jpeg" {
		filePath := helpers.GetFilePath("cake")
		if _, err := os.Stat(filePath + cakeExist.Image); err == nil {
			if err := os.Remove(filePath + cakeExist.Image); err != nil {
				return err
			}
		}
	}

	delete, err := r.db.PrepareContext(ctx, "DELETE FROM cakes WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = delete.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *cakeRepo) UploadImage(ctx *gin.Context, path string) error {
	file := ctx.Request.MultipartForm.File["file"][0]

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		return err
	}
	return nil
}
