package model

import "time"

// Category，菜品的类别，比如川菜、粤菜、热菜、凉菜、饮品
// ID 唯一标识，用于查找
// Name 分类名称，如川菜、湘菜
// Sort 控制在列表中的显示顺序
// Status 表示启用/禁用
// CteatedAt 创建时间
// UpdatedAt 更新时间

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Sort      int       `json:"sort"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
