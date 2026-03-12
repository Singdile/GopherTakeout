package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresDB 创建数据库连接池
// dsn: 数据库连接字符串，如 "host=localhost user=postgres password=123456 dbname=mydb port=5432 sslmode=disable"
func NewPostgresDB(dsn string) (*pgxpool.Pool, error) {
	//创建一个数据库连接池Pool,返回指针，err
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	//Ping触发完整的握手流程，测试是否能够正常连接
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return pool, nil
}
