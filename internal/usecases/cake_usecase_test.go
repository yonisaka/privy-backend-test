package usecases_test

import (
	"errors"
	"privy-backend-test/internal/domain"
	"privy-backend-test/internal/domain/repository"
	"privy-backend-test/internal/domain/usecase"
	"privy-backend-test/internal/usecases"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	InternalServerError = errors.New("internal server error")
)

type fields struct {
	cakeRepo repository.CakeRepository
}

func sut(f fields) usecase.CakeUsecase {
	return usecases.NewCakeUsecase(f.cakeRepo)
}

func TestGetCakes(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}

	type test struct {
		fields  fields
		args    args
		want    *[]domain.Cake
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling cake repo GetCakes succeed, Should return no error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						GetCakesFunc: func(ctx *gin.Context) (*[]domain.Cake, error) {
							assert.Equal(t, args.ctx, ctx)

							return &[]domain.Cake{
								{
									ID:          1,
									Title:       "Cake Test",
									Description: "Cake Test Description",
									Rating:      5.5,
									Image:       "cake-test.jpg",
								},
							}, nil
						},
					},
				},
				args: args,
				want: &[]domain.Cake{
					{
						ID:          1,
						Title:       "Cake Test",
						Description: "Cake Test Description",
						Rating:      5.5,
						Image:       "cake-test.jpg",
					},
				},
				wantErr: nil,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := sut(tt.fields)

			got, err := sut.GetCakes(tt.args.ctx)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetCakeByID(t *testing.T) {
	type args struct {
		ctx *gin.Context
		id  int64
	}

	type test struct {
		fields  fields
		args    args
		want    *domain.Cake
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling cake repo GetCakeByID succeed, Should return no error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				id:  1,
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						GetCakeByIDFunc: func(ctx *gin.Context, id int64) (*domain.Cake, error) {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.id, id)

							return &domain.Cake{
								ID:          1,
								Title:       "Cake Test",
								Description: "Cake Test Description",
								Rating:      5.5,
								Image:       "cake-test.jpg",
							}, nil
						},
					},
				},
				args: args,
				want: &domain.Cake{
					ID:          1,
					Title:       "Cake Test",
					Description: "Cake Test Description",
					Rating:      5.5,
					Image:       "cake-test.jpg",
				},
				wantErr: nil,
			}
		},
		"Given invalid request parameter, When calling cake repo GetCakeByID failed, Should return error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				id:  100,
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						GetCakeByIDFunc: func(ctx *gin.Context, id int64) (*domain.Cake, error) {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.id, id)

							return nil, InternalServerError
						},
					},
				},
				args:    args,
				want:    nil,
				wantErr: InternalServerError,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := sut(tt.fields)

			got, err := sut.GetCakeByID(tt.args.ctx, tt.args.id)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStoreCake(t *testing.T) {
	type args struct {
		ctx *gin.Context
		c   *domain.Cake
	}

	type test struct {
		fields  fields
		args    args
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling cake repo StoreCake succeed, Should return no error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				c: &domain.Cake{
					Title:       "Cake Test",
					Description: "Cake Test Description",
					Rating:      5.5,
					Image:       "cake-test.jpg",
				},
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						StoreFunc: func(ctx *gin.Context, c *domain.Cake) error {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.c, c)

							return nil
						},
					},
				},
				args:    args,
				wantErr: nil,
			}
		},
		"Given invalid request parameter, When calling cake repo StoreCake failed, Should return error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				c: &domain.Cake{
					Title:       "Cake Test",
					Description: "Cake Test Description",
					Rating:      5.5,
					Image:       "cake-test.jpg",
				},
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						StoreFunc: func(ctx *gin.Context, c *domain.Cake) error {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.c, c)

							return InternalServerError
						},
					},
				},
				args:    args,
				wantErr: InternalServerError,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := sut(tt.fields)

			err := sut.Store(tt.args.ctx, tt.args.c)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCake(t *testing.T) {
	type args struct {
		ctx *gin.Context
		c   *domain.Cake
	}

	type test struct {
		fields  fields
		args    args
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling cake repo UpdateCake succeed, Should return no error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				c: &domain.Cake{
					ID:          1,
					Title:       "Cake Test",
					Description: "Cake Test Description",
					Rating:      5.5,
					Image:       "cake-test.jpg",
				},
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						UpdateFunc: func(ctx *gin.Context, c *domain.Cake) error {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.c, c)

							return nil
						},
					},
				},
				args:    args,
				wantErr: nil,
			}
		},
		"Given invalid request parameter, When calling cake repo UpdateCake failed, Should return error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				c: &domain.Cake{
					ID:          1,
					Title:       "Cake Test",
					Description: "Cake Test Description",
					Rating:      5.5,
					Image:       "cake-test.jpg",
				},
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						UpdateFunc: func(ctx *gin.Context, c *domain.Cake) error {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.c, c)

							return InternalServerError
						},
					},
				},
				args:    args,
				wantErr: InternalServerError,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := sut(tt.fields)

			err := sut.Update(tt.args.ctx, tt.args.c)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteCake(t *testing.T) {
	type args struct {
		ctx *gin.Context
		id  int64
	}

	type test struct {
		fields  fields
		args    args
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"Given valid request parameter, When calling cake repo DeleteCake succeed, Should return no error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				id:  1,
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						DeleteFunc: func(ctx *gin.Context, id int64) error {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.id, id)

							return nil
						},
					},
				},
				args:    args,
				wantErr: nil,
			}
		},
		"Given invalid request parameter, When calling cake repo DeleteCake failed, Should return error": func(t *testing.T) test {
			ctx := gin.Context{}

			args := args{
				ctx: &ctx,
				id:  1,
			}

			return test{
				fields: fields{
					cakeRepo: &repository.CakeRepositoryMock{
						DeleteFunc: func(ctx *gin.Context, id int64) error {
							assert.Equal(t, args.ctx, ctx)
							assert.Equal(t, args.id, id)

							return InternalServerError
						},
					},
				},
				args:    args,
				wantErr: InternalServerError,
			}
		},
	}

	for name, testFn := range tests {
		t.Run(name, func(t *testing.T) {
			tt := testFn(t)

			sut := sut(tt.fields)

			err := sut.Delete(tt.args.ctx, tt.args.id)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
