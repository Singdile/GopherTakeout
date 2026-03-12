package test

import (
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/singdile/GopherTakeout/pkg/database"
)

// NewTestDB 创建测试数据库连接池
func NewTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper() //标记为测试辅助函数,如果该函数内部报错，报错信息会指向调用NewTestDB的地方

	//获取环境变量关于数据库DSN
	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=123456 dbname=gopher_takeout_test port=5434 sslmode=disable"
	}

	pool, err := database.NewPostgresDB(dsn)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	//当前测试结束时执行
	t.Cleanup(func() { pool.Close() })
	return pool
}
