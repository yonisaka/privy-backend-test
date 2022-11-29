package repositories

import (
	"os"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/helpers"
	"time"

	"github.com/gin-gonic/gin"
)

type cakeRepo struct {
	*BaseRepo
}

func NewCakeRepo(base *BaseRepo) repository.CakeRepository {
	return &cakeRepo{BaseRepo: base}
}

func (i *cakeRepo) GetCakes(ctx *gin.Context) (*[]domain.Cake, error) {
	rows, err := i.db.QueryContext(ctx, "SELECT * FROM cakes")
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

func (i *cakeRepo) GetCakeByID(ctx *gin.Context, id int64) (*domain.Cake, error) {
	row, err := i.db.QueryContext(ctx, "SELECT * FROM cakes WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	cake := domain.Cake{}
	for row.Next() {
		err := row.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &cake, nil
}

func (i *cakeRepo) Store(ctx *gin.Context, cake *domain.Cake) error {
	timeNow, err := helpers.ConvertToTime(time.Now())
	if err != nil {
		return err
	}
	cake.CreatedAt = &timeNow
	cake.UpdatedAt = &timeNow
	store, err := i.db.PrepareContext(ctx, "INSERT INTO cakes (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
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

func (i *cakeRepo) Update(ctx *gin.Context, cake *domain.Cake) error {
	timeNow, err := helpers.ConvertToTime(time.Now())
	if err != nil {
		return err
	}
	cake.UpdatedAt = &timeNow
	update, err := i.db.PrepareContext(ctx, "UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?")
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

func (i *cakeRepo) Delete(ctx *gin.Context, id int64) error {
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row, err := tx.QueryContext(ctx, "SELECT image FROM cakes WHERE id = ?", id)
	if err != nil {
		return err
	}
	defer row.Close()

	cake := domain.Cake{}
	for row.Next() {
		err := row.Scan(&cake.Image)
		if err != nil {
			return err
		}
	}

	if cake.Image != "" {
		filePath := helpers.GetFilePath("cake")
		if _, err := os.Stat(filePath + cake.Image); err == nil {
			if err := os.Remove(filePath + cake.Image); err != nil {
				return err
			}
		}
	}
	delete, err := tx.PrepareContext(ctx, "DELETE FROM cakes WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = delete.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (i *cakeRepo) UploadImage(ctx *gin.Context, path string) error {
	file := ctx.Request.MultipartForm.File["file"][0]

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		return err
	}
	return nil
}
