# 基于Golang实现的微服务项目

##### 注:个人学习微服务使用

## using the technology stack

1. Golang
2. Mysql
3. Redis
4. Beego
5. Consul
6. Micro
7. GRPC
8. Docker
9. Nginx
10. FastDFS
11. Protobuf

## 功能模块

用户模块

- [x] 注册

  - [x]  获取验证码图片服务
  - [x]  获取短信验证码服务
  - [x]  发送注册信息服务
- [x] 登录

  - [x]  获取session信息服务
  - [x]  获取登录信息服务
- [x] 退出
  - [x] 删除(退出)登录信息服务
- [x] 个人信息获取

  - [x]  获取用户基本信息服务
  - [x]  发送上传用户头像服务
  - [x]  更新用户名服务
- [x] 实名认证

  - [x] 获取用户实名信息服务
- [x] 发送用户实名认证信息服务

房屋模块

- [x] 首页展示
  - [x]  获取首页banner图服务
- [x] 房屋详情

  - [x]  获取用户已发布房源信息服务
  - [x]  发布房屋详细信息的服务
  - [x]  上传房屋图片的服务
  - [x]  获取房屋详细信息服务
- [x] 地区列表
  - [x] 获取地区信息服务
- [x] 房屋搜索 
  - [x] 获取(搜索)房源服务
- [x] 订单模块

  - [x]  发布(发送)订单服务
  - [x]  获取房东/租户订单信息服务
  - [x]  更新房东同意/拒绝订单
  - [x]  更新用户评价订单信息
  - [x]  查看订单信息
  - [x]  订单评论


## project structure

```bash
.
├── DeleteSession
│   ├── 退出登录时清除Session
├── GetArea
│   ├── 获取地区信息服务
├── GetHouseInfo
│   ├── 获取房屋详细信息服务
├── GetHouses
│   ├── 获取（搜索）房源服务
├── GetImageCd
│   ├── 获取验证码图片服务
├── GetIndex
│   ├── 获取首页Banner图片服务
├── GetSession
│   ├── 获取session信息服务
├── GetSmscd
│   ├── 获取短信验证码服务
├── GetUserHouses
│   ├── 获取用户已发布房源信息服务
├── GetUserInfo
│   ├── 获取用户基本信息服务
├── GetUserOrder
│   ├── 获取房东/租户订单信息服务
├── PostAvatar
│   ├── 发送（上传）用户头像服务
├── PostHouses
│   ├── 发送（发布）房源信息服务
├── PostHousesImage
│   ├── 发送（上传）房屋图片服务
├── PostLogin
│   ├── 发送登陆信息服务
├── PostOrders
│   ├── 发送（发布）订单服务
├── PostRet
│   ├── 发送注册信息服务
├── PostUserAuth
│   ├── 发送用户实名认证信息服务
├── PutComment
│   ├── 更新用户评价订单信息
├── PutOrders
│   ├── 更新房东同意/拒绝订单
├── PutUserInfo
│   ├── 更新用户名服务
├── homeWeb
│   ├── conf 项目配置文件
│   │   ├── app.conf
│   │   ├── data.sql
│   │   └── redis.conf
│   ├── handler
│   │   └── handler.go 配置路由
│   ├── html 项目静态文件 index.html
│   ├── main.go 主函数
│   ├── model 数据库模型
│   │   └── models.go
│   ├── plugin.go
│   ├── server.sh
│   └── utils 项目中用到的工具函数
│       ├── config.go
│       ├── error.go
│       └── misc.go
└── README.md
```

## Documentation



## Quick start

> 执行以下命令之前确保你的服务器开发环境已经配置完整.

- consul启动：

  ```bash
  开发测试过程中可以使用单机模式
  consul agent -dev
  ```

- 创建Micro服务

  ```bahs
  micro new --type "web" Microservices/homeWeb
  
  micro new --type "srv" Microservices/GetArea
  # "srv" 是表示当前创建的微服务类型
  # Microservices是相对于GOPATH/src下的文件夹名称 可以根据项目进行设置 
  # GetArea是当前创建的微服务的文件名
  ```

- 启动Mysql数据库

  ```bahs
  systemctl start mysql
  ```

- 启动Redis数据库

  ```bahs
   sudo redis-server /etc/redis/redis.conf
  ```

- FastDFS服务启动

  ```bash
  sudo fdfs_trackerd /etc/fdfs/tracker.conf
  sudo fdfs_storaged /etc/fdfs/storage.conf
  ```

- FastDFS 和Nginx整合

  ```bash
  新建一个nginx配置文件nginx-fdfs.conf
  添加一个server区段:
  
  server {
          listen       80;
          server_name  127.0.0.1;
          location /group1/M00/{
          #root/home/FastDFS/fdfs_storage/data;
          ngx_fastdfs_module;
          }
  }
  
  说明：
  server_name指定本机ip
  
  location /group1/M00/：group1为nginx 服务FastDFS的分组名称，M00是FastDFS自动生成编号，对应store_path0=/home/FastDFS/fdfs_storage，如果FastDFS定义store_path1，这里就是M01
  
  启动nginx
  sudo /usr/local/nginx/sbin/nginx
  重启nginx
  sudo /usr/local/nginx/sbin/nginx -s restart
  ```



