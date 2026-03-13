package repository

import (
	"testing"
	"time"

	"github.com/singdile/GopherTakeout/internal/model"
	"github.com/singdile/GopherTakeout/test"
)

// 创建菜品类别失败
func TestCategory_Create_fail(t *testing.T) {
	//连接数据库
	pool := test.NewTestDB(t)
	//准备测试数据
	category := model.Category{
		Name:      "川菜",
		Sort:      1,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//执行

	//断言
}

// 创建菜品类别成功
func TestCategory_Create_success(t *testing.T) {}
