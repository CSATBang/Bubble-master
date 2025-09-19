# Go Todo API 后端服务


------

## 📖 项目简介

一个使用 Go 语言开发的 RESTful API 服务，为待办事项（Todo）应用提供完整的后端支持。本项目是配合前端项目 **[Bubble](https://github.com/Q1mi/bubble_frontend)** 开发的后端接口实现。

------

## 🎯 功能特性

- ✅ 完整的 **CRUD**（创建、读取、更新、删除）
- ✅ 遵循 **RESTful API** 设计规范
- ✅ **MySQL** 数据库集成
- ✅ 使用 **Gin** Web 框架
- ✅ 使用 **GORM** 进行 ORM 操作
- ✅ 清晰的项目结构
- ✅ 基本的错误处理与日志记录

------

## 🔗 关联项目

- **前端项目**:来自B站Qimi老师的**[Bubble](https://github.com/Q1mi/bubble_frontend)**
- **前端技术栈**: Vue/React + TypeScript + Element UI
- **项目功能**: 现代化的待办事项管理界面

------

## 🏗️ 技术栈

**后端**

- 语言: Go 1.19+
- Web 框架: Gin
- ORM: GORM
- 数据库: MySQL 5.7+
- 驱动: Go-MySQL-Driver

**开发工具**

- Go Modules
- Git
- MySQL Workbench
- Postman（API 测试）

------

## 📦 项目结构

若不想看有项目结构的版本，可以下载单文件版本也就是目录里的main_QWQ.go

```text
go-todo-api/
├── main.go          # 应用入口
├── main_QWQ.go      #单文件版本
├── database/        # 数据库配置
│   └── database.go
├── models/          # 数据模型
│   └── todo.go
├── handlers/        # 请求处理器
│   └── todo_handler.go
├── routers/         # 路由配置
│   └── router.go
├── go.mod           # Go 模块定义
└── README.md        # 项目说明
```

------

## 🚀 快速开始

### 前置要求

- Go 1.19 或更高版本
- MySQL 5.7 或更高版本
- Git

### 1. 克隆项目

```bash
git clone https://github.com/your-username/go-todo-api.git
cd go-todo-api
```

### 2. 配置数据库

```sql
CREATE DATABASE bubble CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 配置数据库连接（修改 `database/database.go` 中的 DSN）

```go
dsn := "root:your_password@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
```

> 提示：在生产环境中建议使用环境变量（例如 `os.Getenv` / `.env`）来管理数据库账号和密码，避免将敏感信息写死在代码里。

### 4. 安装依赖

```bash
go mod download
or
go mod tidy
```

### 5. 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:10090` 启动（默认）

<img width="1195" height="641" alt="image-20250919141726574" src="https://github.com/user-attachments/assets/e877688e-0d97-44f5-b64a-c19c361ab7f4" />


------

## 📡 API 接口

**基础 URL**：

```
http://localhost:10090/v1
```

### 接口列表

| 方法   | 端点      | 描述                        | 状态码 |
| ------ | --------- | --------------------------- | ------ |
| POST   | /todo     | 创建新待办事项              | 201    |
| GET    | /todo     | 获取所有待办事项            | 200    |
| GET    | /todo/:id | 获取指定待办事项            | 200    |
| PUT    | /todo/:id | 更新待办事项（例如 status） | 200    |
| DELETE | /todo/:id | 删除待办事项                | 204    |

### 请求示例 - 创建待办事项

```bash
curl -X POST http://localhost:10090/v1/todo \
  -H "Content-Type: application/json" \
  -d '{"title": "学习Go语言", "status": false}'
```

**响应**

```json
{
  "id": 1,
  "title": "学习Go语言",
  "status": false
}
```

------

## 🧪 测试 API

使用 Postman 或 curl：

```bash
# 获取所有待办事项
curl http://localhost:10090/v1/todo

# 获取单个待办事项
curl http://localhost:10090/v1/todo/1

# 更新待办事项（示例：只更新 status 字段）
curl -X PUT http://localhost:10090/v1/todo/1 \
  -H "Content-Type: application/json" \
  -d '{"status": true}'

# 删除待办事项
curl -X DELETE http://localhost:10090/v1/todo/1
```

------

## 📚 学习资源

该项目适合 Go 初学者学习：

- Go 语言基础语法
- Gin Web 框架使用
- GORM 数据库操作
- RESTful API 设计
- 项目结构组织
- 错误处理和日志

------

## 🙏 致谢

- 感谢 七米老师 前端项目提供的设计灵感
- 感谢 Gin 和 GORM 团队提供的优秀库

------

## 📞 联系方式

- 作者: BastienHyde
- GitHub: [CSATBang](https://github.com/CSATBang)
- 邮箱: c508428101@outlook.com

------

⭐ 如果这个项目对你有帮助，请给它一个 Star！
