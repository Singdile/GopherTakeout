package test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/singdile/GopherTakeout/pkg/config"
	"github.com/singdile/GopherTakeout/pkg/database"
)

// NewTestDB 创建测试数据库连接池
func NewTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper() //标记为测试辅助函数,如果该函数内部报错，报错信息会指向调用NewTestDB的地方

	//设置测试环境下的环境变量
	t.Setenv("APP_ENV", "test")

	//初始化配置
	config.InitConfig()
	pool, err := database.NewPostgresDB(config.AppConfig.Database.DSN)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	//当前测试结束时执行
	t.Cleanup(func() {
		pool.Exec(context.Background(), "TRUNCATE categories RESTART IDENTITY CASCADE;")
		pool.Close()
	})
	return pool
}
