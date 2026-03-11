# GopherTakeout

## 项目概述

GopherTakeout 是一个仿苍穹外卖商家端的 Go 语言学习项目。项目使用测试驱动开发（TDD），从零开始实现商家端的完整功能。

**目标用户**: Go 语言初学者（有 Rust 背景）

## 技术栈

| 组件 | 技术 | 版本 |
|------|------|------|
| 语言 | Go | 1.21+ |
| Web框架 | Gin | v1.9+ |
| 数据库 | PostgreSQL | 15+ |
| 数据库驱动 | pgx/v5 | v5.5+ |
| 配置管理 | viper | v1.18+ |
| 测试框架 | testify | v1.8+ |
| 参数验证 | validator | v10+ |

## 常用命令

```bash
# 开发
make run              # 运行服务
make build            # 构建二进制

# 测试
make test             # 运行所有测试
make test-unit        # 单元测试
make test-integration # 集成测试
make test-cover       # 测试覆盖率报告

# 数据库
make migrate-up       # 执行迁移
make migrate-down     # 回滚迁移
make db-reset         # 重置数据库

# 代码质量
make lint             # 代码检查
make fmt              # 格式化代码
```

## 项目结构

```
GopherTakeout/
├── cmd/                    # 应用入口
│   └── server/main.go     # 服务启动入口
├── internal/              # 私有应用代码（不可被外部导入）
│   ├── handler/          # HTTP处理器（Controller层）
│   ├── service/          # 业务逻辑层
│   ├── repository/       # 数据访问层（DAO）
│   └── model/            # 数据模型/实体
├── pkg/                   # 可复用的公共库
│   ├── config/           # 配置管理
│   ├── database/         # 数据库连接池
│   └── response/         # 统一响应格式
├── test/                  # 集成测试
├── migrations/            # 数据库迁移脚本
├── configs/               # 配置文件
├── Makefile              # 常用命令
└── AGENTS.md             # 本文件
```

## 架构分层

采用经典三层架构：

```
Handler → Service → Repository → Database
```

**职责划分**:

- **Handler**: 处理HTTP请求/响应，参数验证
- **Service**: 业务逻辑，事务管理
- **Repository**: 数据库CRUD操作
- **Model**: 数据结构定义

## 开发规范

### 命名规范

- 包名: 小写单词，不使用下划线
- 文件名: 小写+下划线，如 `category_service.go`
- 接口: 动词+名词，如 `CategoryRepository`
- 结构体: 名词，如 `Category`
- 函数: 驼峰命名，导出函数首字母大写

### 错误处理

```go
// 定义业务错误
var (
    ErrNotFound      = errors.New("resource not found")
    ErrAlreadyExists = errors.New("resource already exists")
)

// 错误包装
if err != nil {
    return fmt.Errorf("failed to get category: %w", err)
}
```

### 测试规范

```go
// 文件命名: xxx_test.go
// 函数命名: Test<功能>_<场景>

func TestCategoryRepository_Create_Success(t *testing.T) {
    // Arrange
    // Act
    // Assert
}
```

**TDD流程**:
1. 编写失败的测试
2. 编写最小代码使测试通过
3. 重构代码

### 响应格式

```go
// 成功
{
    "code": 200,
    "message": "success",
    "data": {...}
}

// 失败
{
    "code": 400,
    "message": "invalid parameter",
    "data": null
}
```

## 数据库Schema

### 分类表 (categories)
```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    sort INT DEFAULT 0,
    status SMALLINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 菜品表 (dishes)
```sql
CREATE TABLE dishes (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id),
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    description TEXT,
    image VARCHAR(255),
    status SMALLINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 店铺表 (shops)
```sql
CREATE TABLE shops (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255),
    phone VARCHAR(20),
    status SMALLINT DEFAULT 1,
    opening_time TIME,
    closing_time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 订单表 (orders)
```sql
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    shop_id INT REFERENCES shops(id),
    order_number VARCHAR(50) NOT NULL,
    status SMALLINT DEFAULT 0,
    total_amount DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## API设计

### 分类管理 (Category)

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/categories | 获取分类列表 |
| GET | /api/v1/categories/:id | 获取分类详情 |
| POST | /api/v1/categories | 创建分类 |
| PUT | /api/v1/categories/:id | 更新分类 |
| DELETE | /api/v1/categories/:id | 删除分类 |

### 菜品管理 (Dish)

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/dishes | 获取菜品列表 |
| GET | /api/v1/dishes/:id | 获取菜品详情 |
| POST | /api/v1/dishes | 创建菜品 |
| PUT | /api/v1/dishes/:id | 更新菜品 |
| DELETE | /api/v1/dishes/:id | 删除菜品 |

### 店铺管理 (Shop)

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/shops/:id | 获取店铺信息 |
| PUT | /api/v1/shops/:id | 更新店铺信息 |
| PUT | /api/v1/shops/:id/status | 更新营业状态 |

### 订单管理 (Order)

| Method | Path | Description |
|--------|------|-------------|
| GET | /api/v1/orders | 获取订单列表 |
| GET | /api/v1/orders/:id | 获取订单详情 |
| PUT | /api/v1/orders/:id/status | 更新订单状态 |

## 学习路径

按模块顺序学习：

1. **Model层** → struct、tag、json序列化
2. **Repository层** → pgx、SQL、interface
3. **Service层** → 错误处理、业务逻辑
4. **Handler层** → Gin框架、HTTP处理
5. **集成测试** → httptest、测试数据库

## Go vs Rust 对照

| 概念 | Go | Rust |
|------|-----|------|
| 结构体 | struct + tag | struct + derive |
| 错误处理 | error接口 | Result<T, E> |
| 接口 | 隐式实现 | trait |
| 并发 | goroutine + channel | async/await |
| 包管理 | go mod | cargo |
| 测试 | testing包 | #[test] |

## 注意事项

- Go 没有 class 和继承，使用 struct 和 interface
- 错误处理使用返回值，不是异常
- 包的可见性通过首字母大小写控制
- 没有泛型约束（Go 1.18+有泛型，但建议先学基础）
- defer 语句用于资源清理（类似 Rust 的 Drop）