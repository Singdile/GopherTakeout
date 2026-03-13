package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/singdile/GopherTakeout/internal/model"
)

type CategoryRepository interface {
	//查找所有菜品类型
	FindAll(ctx context.Context, pool *pgxpool.Pool) ([]model.Category, error)

	//通过ID查找菜品类型
	FindByID(ctx context.Context, id int, pool *pgxpool.Pool) (*model.Category, error)

	//创造新的菜品类型
	Create(ctx context.Context, category *model.Category, pool *pgxpool.Pool) error

	//更新菜品类型
	Update(ctx context.Context, category *model.Category, pool *pgxpool.Pool) error

	//删除菜品类型
	Delete(ctx context.Context, id int, pool *pgxpool.Pool) error
}
