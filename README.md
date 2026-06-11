# DataHub - 测试开发学习平台

一个基于 Go + Vue3 的测试开发学习平台，用于学习和实践 API 测试、自动化测试等技能。

## 功能特性

- **用户管理**：登录/注册系统，支持 Mock 数据测试
- **测试用例管理**：创建、编辑、删除、查询测试用例
- **API测试执行**：支持 HTTP/HTTPS 请求，支持 GET/POST/PUT/DELETE/PATCH 方法
- **测试报告**：查看测试执行结果和历史记录
- **深色主题**：现代化的玻璃拟态设计风格
- **响应式布局**：支持桌面端和移动端

## 技术栈

### 后端
- Go 1.20+
- Gin Web Framework
- SQLite（嵌入式数据库）
- JWT 认证

### 前端
- Vue 3 + Vite
- Naive UI 组件库
- Vue Router
- Axios

## 快速开始

### 前端运行

```bash
cd DataHub/frontend
npm install
npm run dev
```

访问：http://localhost:5173/

### 后端运行

```bash
cd DataHub/backend
go run main.go
```

后端服务：http://localhost:8080/

## 项目结构

```
DataHub/
├── backend/                    # Go后端代码
│   ├── config/                # 配置文件
│   ├── controllers/           # API控制器
│   │   ├── user_controller.go     # 用户相关接口
│   │   ├── testcase_controller.go # 测试用例接口
│   │   └── testexec_controller.go # 测试执行接口
│   ├── middleware/            # 中间件
│   │   └── auth.go                # JWT认证中间件
│   ├── models/                # 数据模型
│   │   ├── user.go                # 用户模型
│   │   ├── testcase.go            # 测试用例模型
│   │   └── testreport.go          # 测试报告模型
│   ├── routes/                # 路由配置
│   │   └── routes.go              # API路由定义
│   ├── utils/                 # 工具函数
│   │   ├── db.go                  # 数据库连接
│   │   └── jwt.go                 # JWT工具
│   ├── go.mod
│   ├── go.sum
│   └── main.go                # 入口文件
└── frontend/                  # Vue3前端代码
    ├── src/
    │   ├── components/        # 组件
    │   │   ├── Layout.vue         # 布局组件
    │   │   └── Sidebar.vue        # 侧边栏组件
    │   ├── router/            # 路由配置
    │   │   └── index.js           # Vue Router配置
    │   ├── utils/             # 工具函数
    │   │   └── axios.js           # Axios封装与Mock数据
    │   ├── views/             # 页面视图
    │   │   ├── Login.vue          # 登录页面
    │   │   ├── Register.vue       # 注册页面
    │   │   ├── Dashboard.vue      # 仪表盘页面
    │   │   ├── TestCases.vue      # 测试用例列表
    │   │   ├── CreateTestCase.vue # 新建测试用例
    │   │   ├── EditTestCase.vue   # 编辑测试用例
    │   │   ├── Reports.vue        # 报告列表
    │   │   └── ReportDetail.vue   # 报告详情
    │   ├── App.vue            # 根组件
    │   ├── main.js            # 入口文件
    │   └── style.css          # 全局样式
    ├── index.html
    ├── package.json
    └── vite.config.js         # Vite配置
```

## API接口

### 用户认证

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/auth/login | 用户登录 |
| POST | /api/auth/register | 用户注册 |
| GET | /api/profile | 获取用户信息 |

### 测试用例

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/testcases | 创建测试用例 |
| GET | /api/testcases | 获取测试用例列表 |
| GET | /api/testcases/:id | 获取单个测试用例 |
| PUT | /api/testcases/:id | 更新测试用例 |
| DELETE | /api/testcases/:id | 删除测试用例 |

### 测试执行

| 方法 | 路径 | 描述 |
|------|------|------|
| POST | /api/exec/testcase/:id | 执行测试用例 |
| GET | /api/reports | 获取测试报告列表 |
| GET | /api/reports/:id | 获取测试报告详情 |

## Mock测试数据

### 测试账号

- **用户名**: `testuser`
- **密码**: `123456`

### 测试用例示例

系统预置了3个示例测试用例，可直接用于测试平台功能。

## 登录页面

项目包含两个登录页面：

1. **标准登录页** (`/login`)：深色主题，动态粒子背景
2. **3D沉浸感登录页** (`Login3D.html`)：Three.js 3D几何体背景，毛玻璃卡片效果

## License

MIT License
