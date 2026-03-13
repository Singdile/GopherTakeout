package repository

import (
	"context"
	"testing"
	"time"

	"github.com/singdile/GopherTakeout/internal/model"
	"github.com/singdile/GopherTakeout/test"
)

// 创建菜品类别成功
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

// 按照ID查询菜品类型成功
func TestFindByID_Success(t *testing.T) {
	//连接数据库
	pool := test.NewTestDB(t)
	categoryRepository := NewCategoryRepository(pool)

	category := model.Category{
		Name:   "川菜",
		Sort:   1,
		Status: 1,
	}
	//准备测试数据
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //使用context工作机制，超过5s，关闭资源，返回超时错误
	defer cancel()
	categoryRepository.Create(ctx, &category)
	var id int = 1

	//执行
	_, err := categoryRepository.FindByID(ctx, id)

	//断言
	if err != nil {
		t.Errorf("Find by id fail,id: %d, error: %s", id, err.Error())
	}
}
