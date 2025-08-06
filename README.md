# Empire Origins - Go 版本

<p align="center">
  <img src="https://go.dev/images/go-logo-blue.svg" width="120" alt="Go Logo" />
</p>

[![Go Version](https://img.shields.io/badge/Go-1.24.5+-blue.svg)](https://golang.org/)
[![Gin Version](https://img.shields.io/badge/Gin-1.10.1+-green.svg)](https://gin-gonic.com/)
[![GORM Version](https://img.shields.io/badge/GORM-1.30.1+-orange.svg)](https://gorm.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 项目描述

Empire Origins 是一个基于 Go 语言开发的战国策略类游戏，这是对原有 Node.js 后端版本的 Go 语言重写实现。项目采用现代化的 Go 技术栈，提供高性能、可扩展的游戏服务器解决方案。

### 主要特性

- 🚀 **高性能**: 基于 Go 语言的并发特性，提供高并发处理能力
- 🏗️ **模块化架构**: 采用清晰的分层架构，便于维护和扩展
- 🗄️ **数据库支持**: 集成 PostgreSQL 和 Redis，支持复杂的数据存储需求
- 🔐 **安全认证**: 集成 JWT 认证机制
- 📊 **API 设计**: RESTful API 设计，支持 GraphQL（计划中）
- 🐳 **容器化**: 支持 Docker 部署
- 📝 **日志系统**: 完整的日志记录和监控

## 技术栈

### 核心框架

- **Go 1.24.5+**: 主要开发语言
- **Gin**: 高性能 HTTP Web 框架
- **GORM**: Go 语言的 ORM 库

### 数据库

- **PostgreSQL**: 主数据库
- **Redis**: 缓存和会话存储

### 开发工具

- **godotenv**: 环境变量管理
- **validator**: 数据验证

## 项目结构

```
empire_origins_golang/
├── config/           # 配置管理
│   └── config.go
├── db/              # 数据库连接
│   ├── database.go
│   └── redis.go
├── middleware/      # 中间件
├── model/          # 数据模型
│   └── user.go
├── router/         # 路由定义
│   └── app.router.go
├── service/        # 业务逻辑
│   └── user.service.go
├── go.mod          # Go 模块定义
├── go.sum          # 依赖校验
├── main.go         # 应用入口
└── README.md       # 项目文档
```

## 快速开始

### 环境要求

- Go 1.24.5 或更高版本
- PostgreSQL 数据库
- Redis 服务器

### 安装依赖

```bash
# 安装 Go 依赖
go mod download
```

### 环境配置

创建 `.env` 文件并配置以下环境变量：

```env
# 服务器配置
PORT=6001

# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=empire_origins

# Redis 配置
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

### 运行项目

```bash
# 开发模式
go run main.go

# 构建并运行
go build -o empire-origins
./empire-origins
```

## API 文档

### 用户相关 API

#### 获取用户信息

```
GET /user/info?id={user_id}
```

**响应示例:**

```json
{
  "id": "1",
  "username": "player1",
  "email": "player1@example.com",
  "order": 1,
  "is_enabled": true,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## 开发指南

### 添加新的 API 端点

1. 在 `service/` 目录下创建业务逻辑
2. 在 `router/app.router.go` 中添加路由
3. 在 `model/` 目录下定义数据模型（如需要）

### 数据库迁移

项目使用 GORM 进行数据库操作，支持自动迁移：

```go
// 在 main.go 中添加自动迁移
db.DB.AutoMigrate(&model.User{})
```

### 测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./service

# 运行测试并显示覆盖率
go test -cover ./...
```

## 部署

### Docker 部署

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o empire-origins .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/empire-origins .
EXPOSE 6001
CMD ["./empire-origins"]
```

### 使用 Docker Compose

```yaml
version: "3.8"
services:
  app:
    build: .
    ports:
      - "6001:6001"
    environment:
      - PORT=6001
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: empire_origins
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: your_password
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    volumes:
      - redis_data:/data

volumes:
  postgres_data:
  redis_data:
```

## 与 Node.js 版本的对比

| 特性     | Node.js 版本        | Go 版本           |
| -------- | ------------------- | ----------------- |
| 框架     | NestJS              | Gin               |
| 数据库   | Prisma + PostgreSQL | GORM + PostgreSQL |
| 缓存     | Redis               | Redis             |
| 消息队列 | RabbitMQ            | 计划中            |
| GraphQL  | Apollo Server       | 计划中            |
| 认证     | JWT                 | JWT               |
| 容器化   | Docker              | Docker            |

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。
