# go_admin_project_40

## 教务管理系统（Go + Gin + GORM）

这是一个由 Go 语言实现的简易教务管理系统，作为课程项目 / 作业示例，用于演示学生管理功能 —— 包括学生信息的增删查等。

### 技术栈

- Go  
- Web 框架：Gin  
- ORM：GORM  
- 数据库：SQLite（`school.db`，轻量级，开箱即用）  

### 功能

- 添加学生信息（姓名 / 年龄 / 班级 / 性别）  
- 查询所有学生信息  

### 项目结构
```
go_admin_project_40/
├── cmd/
│ └── server/ ← 服务启动入口
│ └── main.go
├── internal/
│ ├── model/ ← 数据模型定义
│ │ └── student.go
│ ├── repository/ ← 数据访问层
│ │ └── student_repo.go
│ └── service/ ← 业务逻辑层
│ └── student_service.go
└── go.mod / go.sum (如有) ← Go 模块文件
```

快速上手
1. 克隆仓库

首先，克隆项目仓库：
```
git clone https://github.com/xianghuanwwdn/go_admin_project_40.git
cd go_admin_project_40
```
2. 安装依赖

在项目目录中，安装 Go 依赖：
```
go mod tidy
```
3. 启动服务

运行服务：
```
go run cmd/server/main.go
```

服务启动后，监听端口 :8080，你可以访问以下 API。

测试 API
添加学生

请求方法：POST
```
URL：http://localhost:8080/students
```
请求体：
```
{
  "name": "张三",
  "age": 21,
  "class": "计算机科学",
  "gender": "男"
}
```
查询学生

请求方法：GET
```
URL：http://localhost:8080/students
```
