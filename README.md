# 🚀 个人博客系统后端

> 基于 Go + Gin + GORM 开发的个人博客系统后端API服务

## 📋 项目简介

这是一个使用 Go 语言开发的个人博客系统后端，采用 Gin 框架和 GORM ORM 库，实现了完整的博客文章管理、用户认证和评论功能。

### ✨ 核心特性

- 🔐 **JWT 用户认证** - 安全的用户注册、登录和权限管理
- 📝 **文章管理** - 完整的 CRUD 操作，支持文章创建、编辑、删除
- 💬 **评论系统** - 用户可以对文章发表评论
- 🗄️ **数据库支持** - 支持 MySQL 数据库，使用 GORM ORM
- ⚡ **高性能** - 基于 Go 语言和 Gin 框架的高性能 API
- 🔒 **安全性** - 密码加密、JWT Token、输入验证

## 🏗️ 项目架构

```
Golang_task4/
├── 📁 CONF/              # 配置管理
│   ├── config.go         # 配置加载逻辑
│   ├── config.json       # 配置文件
│   └── jwt.go            # JWT 相关配置
├── 📁 DAO/               # 数据访问层
│   ├── db.go             # 数据模型定义
│   ├── dblink.go         # 数据库连接
│   ├── user.go           # 用户数据模型
│   ├── post.go           # 文章数据模型
│   └── comment.go        # 评论数据模型
├── 📁 SERVICE/           # 业务逻辑层
│   ├── user.go           # 用户业务逻辑
│   ├── post.go           # 文章业务逻辑
│   └── comment.go        # 评论业务逻辑
├── 📁 HANDLER/           # 控制器层
│   ├── user.go           # 用户控制器
│   ├── post.go           # 文章控制器
│   └── comment.go        # 评论控制器
├── 📁 ROUTE/             # 路由层
│   ├── user.go           # 用户路由
│   ├── post.go           # 文章路由
│   └── comment.go        # 评论路由
├── 📁 MIDDLEWARE/        # 中间件
│   └── is_login.go       # 登录验证中间件
├── 📁 TEST/              # 测试相关
│   ├── testdata.md       # API 测试文档
│   └── *.png             # 测试截图
├── 📁 CMD/               # 命令行工具（可选）
│   ├── main.go           # CMD 入口
│   └── dblink.go         # CMD 数据库连接
├── 📄 main.go            # 应用程序入口
├── 📄 go.mod             # Go 模块定义
├── 📄 go.sum             # 依赖锁定文件
└── 📄 .gitignore         # Git 忽略文件
```

## 🚀 快速开始

### 环境要求

- **Go**: 1.18+
- **MySQL**: 5.7+
- **Git**: 版本控制

### 1. 克隆项目

```bash
git clone <repository-url>
cd Golang_task4
```

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 数据库配置

1. 创建 MySQL 数据库：
```sql
CREATE DATABASE golang_test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改配置文件 `CONF/config.json`：
```json
{
    "jwt_secret": "your-secret-key-here",
    "dsn": "username:password@tcp(127.0.0.1:3306)/golang_test?charset=utf8mb4&parseTime=True&loc=Local"
}
```

### 4. 启动服务

```bash
go run main.go
```

服务将在 `http://127.0.0.1:8080` 启动。

## 📚 API 文档

### 用户管理

#### 用户注册
```http
POST /user/register
Content-Type: application/json

{
    "username": "testuser",
    "email": "test@example.com",
    "password": "Test123!"
}
```

#### 用户登录
```http
POST /user/login
Content-Type: application/json

{
    "username": "testuser",
    "password": "Test123!"
}
```

响应包含 JWT Token，需要在后续请求的 Header 中携带：
```http
Authorization: Bearer <your-jwt-token>
```

### 文章管理

#### 创建文章
```http
POST /post/add
Authorization: Bearer <token>
Content-Type: application/json

{
    "title": "文章标题",
    "content": "文章内容"
}
```

#### 获取文章列表
```http
GET /post/queryList
```

#### 获取单篇文章
```http
GET /post/query/:id
```

#### 编辑文章
```http
POST /post/edit
Authorization: Bearer <token>
Content-Type: application/json

{
    "id": 1,
    "title": "新标题",
    "content": "新内容"
}
```

#### 删除文章
```http
POST /post/delete
Authorization: Bearer <token>
Content-Type: application/json

{
    "id": 1
}
```

### 评论管理

#### 发表评论
```http
POST /comment/add
Authorization: Bearer <token>
Content-Type: application/json

{
    "post_id": 1,
    "content": "评论内容"
}
```

#### 获取文章评论
```http
GET /comment/query/:post_id
```

## 🔧 开发指南

### 项目结构说明

- **CONF**: 配置管理，使用 Viper 加载配置文件
- **DAO**: 数据访问层，定义数据模型和数据库操作
- **SERVICE**: 业务逻辑层，处理核心业务逻辑
- **HANDLER**: 控制器层，处理 HTTP 请求和响应
- **ROUTE**: 路由层，定义 API 路由
- **MIDDLEWARE**: 中间件，处理认证和跨域等

### 添加新功能

1. **定义数据模型** - 在 `DAO/` 目录添加新的模型
2. **实现业务逻辑** - 在 `SERVICE/` 目录添加业务逻辑
3. **创建控制器** - 在 `HANDLER/` 目录添加控制器
4. **配置路由** - 在 `ROUTE/` 目录添加路由

### 数据库迁移

项目启动时会自动创建所需的数据表：
```go
// 在 main.go 中自动执行
DAO.TableAutoMigrate()
```

## 🧪 测试

### API 测试

详细的 API 测试用例和示例请参考：
- [TEST/testdata.md](TEST/testdata.md) - 完整的 API 测试文档

### 运行测试

```bash
# 启动服务后进行测试
go run main.go

# 使用 curl 或 Postman 测试 API
curl -X POST http://127.0.0.1:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"Test123!"}'
```

## 🔒 安全特性

- **密码加密**: 使用 bcrypt 对密码进行加密存储
- **JWT 认证**: 基于 JWT 的无状态认证机制
- **输入验证**: Gin 框架的绑定和验证
- **SQL 注入防护**: 使用 GORM 预编译语句

## 📦 依赖说明

主要依赖包：
- `github.com/gin-gonic/gin` - Web 框架
- `gorm.io/gorm` - ORM 库
- `gorm.io/driver/mysql` - MySQL 驱动
- `github.com/spf13/viper` - 配置管理
- `github.com/dgrijalva/jwt-go` - JWT 处理
- `golang.org/x/crypto/bcrypt` - 密码加密

## 🐛 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查 MySQL 服务是否启动
   - 验证 config.json 中的 DSN 配置
   - 确认数据库是否存在

2. **端口被占用**
   - 修改 main.go 中的端口号
   - 杀死占用端口的进程

3. **依赖安装失败**
   - 检查网络连接
   - 尝试设置 GOPROXY: `go env -w GOPROXY=https://goproxy.cn,direct`

## 🤝 贡献指南

1. Fork 本项目
2. 创建特性分支: `git checkout -b feature/AmazingFeature`
3. 提交更改: `git commit -m 'Add some AmazingFeature'`
4. 推送分支: `git push origin feature/AmazingFeature`
5. 提交 Pull Request

