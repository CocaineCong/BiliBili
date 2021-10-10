# FiliFili Fan站视频网站 

**此项目使用`Gin`+`Gorm` ，基于`RESTful API`实现的一个B站**。

**此项目比较适合小白进阶`web开发`这方面**

# 接口文档

[BiliBili 接口文档](https://www.showdoc.cc/1621442994395086)

**密码：0000**

![在这里插入图片描述](https://img-blog.csdnimg.cn/ca00b91683434f38bc940879898453c3.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA5bCP55Sf5Yeh5LiA,size_20,color_FFFFFF,t_70,g_se,x_16)

# 项目主要功能介绍

- 用户模块：
    - 注册登录
    - 修改个人信息，更换头像，更改密码
    - 关注用户
    - 粉丝列表
    
   
- 视频模块：
    - 个人上传视频，可加视频封面
    - 点赞，收藏，转发视频
    - 更新视频简介，封面
    - 查看收藏视频列表
    - 删除视频

- 弹幕模块：
    - 发送弹幕
    - 获取弹幕

- 评论模块：
    - 评论他人
    - 回复他人

## 项目主要依赖：

**Golang V1.16**

- Gin
- Gorm
- mysql
- redis
- viper
- jwt-go
- cron
- qiniu-go-sdk

## 项目结构

```shell
BiliBili/
├── api
├── cache
├── conf
├── middleware
├── model
├── pkg
│  ├── e
│  ├── util
├── routes
├── serializer
└── servive
```

- api : 用于定义接口函数
- cache : 放置redis缓存
- conf : 用于存储配置文件
- middleware : 应用中间件
- model : 应用数据库模型
- pkg/e : 封装错误码
- pkg/util : 工具函数
- routes : 路由逻辑处理
- serializer : 将数据序列化为 json 的函数
- servive : 接口函数的实现

## 配置文件

**conf/config.ini**
```yml
server:
  port: 3000
  version: 1.0
  coding: mp4
  jwtSecret: something-very-secret
  adminJwtSecret: admin-secret
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: bilibili
  username: root
  password: root
  charset: utf8mb4
qiniu:
  AccessKey: 
  SerectKey: 
  Bucket: 
  QiniuServer: 
#email:
#  port: 465
#  host: smtp.163.com
#  address: 邮箱地址
#  password: 邮箱授权码
redis:
  address: 127.0.0.1:6379
  password:
admin:
  email: admin@qq.com
  password: admin
```

## 简要说明
1. `mysql`是存储主要数据
2. `redis`用来存储点赞，收藏，浏览这些高实时的


## 项目运行

**本项目使用`Go Mod`管理依赖。**

**下载依赖**

```shell
go mod tidy
```

**运行**

```shell
go run main.go
```

