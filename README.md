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

### 项目部分逻辑
#### 登录的实现

与用户的登录注册有关的接口有三个/douyin/user、/douyin/user/register/、/douyin/user/login其中user主要存储用户信息，不必多言。这里主要关注一下登录和注册功能的具体实现。

##### /douyin/user/register/用户注册

用户注册接口的实现大致逻辑为：

获取用户输入的username和password之后，首先对其进行长度验证（不可超过32位），然后使用Gorm链接数据库，在此期间若遇到错误会有相应的文字错误提示。然后生成useriID，userID是从100000开始，以后每新增一个用户向后递增1，在这里主要是参考了qq账号和B站uid的位数，id仅是账户唯一标识的功能。最后生成token，这里使用了JWT生成token，因为对于单体应用而言，HS256 和 RS256 的安全性没有任何差别。当以上三步都完毕之后，返回信息。

##### /douyin/user/login/用户登录

用户登录接口的实现大致逻辑为：

获取用户输入的username和password，使用gorm的where语句在数据库中对符合该账号和密码的ID进行查找，查找到的话即视为找到对应账户。未找到则提示账户不存在。通过与上方类似的方法生成token。返回信息。

##### 最后

由于使用了gorm，注册和登录时的SQL语句都是防止SQL注入的，安全性可以得到保证。在生成token的时候使用了HS256算法，由于不是微服务架构且是一个手机应用，RS256与HS256安全性相似，可以提高效率。

#### 视频上传部分的逻辑 publish action

1 首先使用过滤器完成token的认证，认证成功进入发布视频的接口，失败就返回认证失败。

2 用户通过再前端app发起请求上传文件，后台通过调用oss文件服务器的上传功能将视频上传到服务器并返回一个url

3 服务器拿到url后调用ffmpeg使用抽取出视频的封面存储到本地

4 服务器再次调用oss的文件上传功能完成封面图的上传，并返回封面图的url

5 服务器收集用户的信息和所发布视频得到的视频url和封面图url，调用gorm的增加接口完成数据库的信息添加

#### 亮点

1 视频存储部分我们使用了oss文件服务器，减轻了本地服务器的压力，本地的服务器不需要大规模的存储文件，数据库也只需要存储视频和封面图的路径。

2 将数据库的存储和封面图的上传这两个操作使用异步来优化，通过使用协程降低了视频发布功能的请求时间，提高了用户的体验。

3 将oss的初始化和bucket的初始化放在项目启动的时候去做，共享一个bucket，减少了bucket的创建的时间，提高了接口的响应速度。

#### 获取用户发布列表的接口 publish list

1 list接口不需要token的拦截器，请求直接放行

2 用户通过再前端app发起请求，后台通过token解析出用户的id

3 服务器通过用户的id查询用户的基本信息和用户的视频列表

4 将查询到的数据进行返回

#### 点赞模块
1. 判断是否是登录状态
2. 下面是查redis的逻辑
   1. 首先先查看当前用户是否对该视频进行了点赞，通过key查询出来用户的点赞信息，会出现三种情况（0,1,2）
      1. 0代表没有点赞呢，1代表点赞了，2代表取消了点赞
   2. 此时判断当前传入的点赞类型和查出的类型
      1. 如果传入的是1，而查出来的不是1，那么就说明，当前用户对该视频点赞成功
      2. 如果传入的是2，而查出来的是1，表示已经点赞了，现在取消了赞，如果查出来的是0或者2，那么什么也不做
   3. 此时应该记录好了用户对某个视频的点赞类型，同时应该也记录上哪个用户给哪个视频点赞了

##### 查询点赞列表
1. 通过查询Set数据结构，传入用户id就可拿到当前用户所有的点赞视频id
2. 然后根据视频id查询出所有的视频

##### 视频点赞总数
通过查询ZSet来查询分数是1的元素数量，也就是统计每个视频的点赞数量

### 项目作者
- https://github.com/code-on-bush
- https://github.com/LiuRuohui
### Git分支说明
- 主分支为完整项目
- 其他分支分别为部分功能模块（可能存在bug）

### 后续计划
- 升级为微服务项目，将功能模块抽取
