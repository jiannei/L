# L

## Gin

* 初始化项目

```
mkdir L && cd L && go mod init L
```

* gin 安装

```
go get -u github.com/gin-gonic/gin
```

* gin 启动

```
go run main.go
```

## 规划

### 目录结构

```
├── README.md               
├── app                     # 应用程序
│   ├── exceptions          # 异常处理
│   ├── http                # Http 服务
│   │   ├── controller      # controller 层
│   │   └── middlewares     # Http 中间件
│   ├── models              # model 层
│   └── providers           # 第三方服务扩展
├── config                  # 应用配置
├── database                # 数据库结构
├── go.mod
├── go.sum
├── main.go                 # 入口文件
├── routes                  # 路由配置
├── storage                 # 存储目录
│   ├── app                 # 规划中...
│   ├── framework           # 规划中...
│   └── logs                # 日志文件
└── tests                   # 单元测试
```

### 功能

- [X] 路由管理 
- [X] 优雅关闭服务器
- [X] 中间件配置
- [] 配置管理
- [] 日志管理
- [] request 请求处理
- [] validator 请求验证器 
- [] response 响应处理
- [] ORM 数据模型
- [] 缓存
- [] RPC 微服务
