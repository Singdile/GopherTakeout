package repository

import (
	"context"
	"testing"
	"time"

	"github.com/singdile/GopherTakeout/internal/model"
	"github.com/singdile/GopherTakeout/test"
)

// 创建菜品类别失败
func TestCategory_Create_Success(t *testing.T) {
	//连接数据库
	pool := test.NewTestDB(t)
	categoryRepository := NewCategoryRepository(pool)

	//准备测试数据
	category := model.Category{
		Name:   "川菜",
		Sort:   1,
		Status: 1,
	}

	//执行
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := categoryRepository.Create(ctx, &category)

	//断言
	if err != nil {
		t.Fatalf("Create a new catogory failed.got error: %v", err)
	}
}

// 插入失败，名称为空
func TestCategory_Create_fail(t *testing.T) {
	//连接数据库
	pool := test.NewTestDB(t)
	categoryRepository := NewCategoryRepository(pool)

	//准备测试数据
	category := model.Category{
		Name:   "",
		Sort:   2,
		Status: 2,
	}

	//执行
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := categoryRepository.Create(ctx, &category)

	//断言
	if err == nil {
		t.Error("Expected error for empty name,got nil.")
	}
}
