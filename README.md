# bytedance-qingxunying
字节跳动青训营抖音项目

## 目录

### 项目架构
本项目采用单体架构，主要以Gin框架为主，采用MVC分层思想降低模块与模块之间的耦合度

### 项目技术栈

- GIN
- Viper
- Zap
- GORM
- MySQL
- Redis
- 阿里云对象存储OOS
- ffmpeg

### 项目基础环境
- Go 1.18
- MySQL 5.7
- Redis 6.x
- Gorm 2.x

### 文件目录
```go
├── common (用于存放业务状态码以及请求参数和响应参数)
│   ├── code.go
│   ├── request.go
│   └── response.go
├── conf  (配置文件信息)
│   └── config.yml
├── controller (对外提供给客户端的接口层)
│   ├── comment.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── dao  (数据库操作层)
│   ├── mysql (对mysql的操作)
│   │   ├── comment.go
│   │   ├── feed.go
│   │   ├── mysql.go
│   │   ├── registInfo.go
│   │   └── user.go
│   └── redis (对redis的操作)
│       ├── consts.go
│       ├── favorite.go
│       ├── redis.go
│       ├── relation.go
├── logger (日志管理)
│   └── logger.go
├── middleware (中间件)
│   └── tokenMiddlewart.go
├── model (模型层)
│   ├── comment.go
│   ├── feed.go
│   ├── registInfo.go
│   └── user.go
├── router (路由)
│   └── router.go
├── service (业务逻辑层)
│   ├── comment.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── setting (viper管理文件)
│   └── setting.go
├── util (工具类)
│   ├── OSSUtil.go
│   └── TokenUtil.go
└── web_app.log (生成的日志文件)
└── main.go (主启动文件)
```

### 项目作者
- https://github.com/code-on-bush
- https://github.com/LiuRuohui
### Git分支说明
- 主分支为完整项目
- 其他分支分别为部分功能模块（可能存在bug）

### 后续计划
- 升级为微服务项目，将功能模块抽取
