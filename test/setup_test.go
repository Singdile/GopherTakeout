package test

import (
	"testing"
)

func TestDB_Connect(t *testing.T) {
	//创建测试数据库连接
	pool := NewTestDB(t)

	//断言
	if pool == nil {
		t.Fatal("expected pool to be non-nil")
	}
}
