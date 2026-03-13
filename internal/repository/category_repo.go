package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/singdile/GopherTakeout/internal/model"
)

// TODO: 实现接口，编写Create的测试
// 实现Category接口的类型，用于操作Model定义的数据
// 这里是用于操作，定义的数据Category
type categoryRepository struct {
	db *pgxpool.Pool
}

// 构造categoryRepository实例
func NewCategoryRepository(db *pgxpool.Pool) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

// 增添新的菜品类型
func (c *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	//验证名称是否为空
	if category.Name == "" {
		return errors.New("Category name cannot be empty") //返回错误信息
	}

	_, err := c.db.Exec(ctx, "INSERT INTO categories (name,sort,status) VALUES ($1,$2,$3) RETURNING id", category.Name, category.Sort, category.Status)

	return err
}
