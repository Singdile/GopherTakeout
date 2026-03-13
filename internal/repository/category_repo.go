package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/singdile/GopherTakeout/internal/model"
	"github.com/singdile/GopherTakeout/pkg/database"
)

// TODO: 实现接口，编写Create的测试
// 实现Category接口的类型，用于操作Model定义的数据
// 这里是用于操作，定义的数据Category
type categoryRepository struct {
	db *pgxpool.Pool
}

// 构造categoryRepository实例
func (c *categoryRepository) NewCategoryRepository() categoryRepository {
	//建立数据连接
	pool, err := database.NewPostgresDB()
}

func (c *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	_, err := c.db.Exec(ctx, "INSERT INTO categories (name,sort,status,createdat,updatedat) VALUES ($1,$2,$3,$4,$5) RETURNING id", category.Name, category.Sort, category.Status, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return err
	}
	return err
}
