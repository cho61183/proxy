# proxy
数据中间件的代理.

# 主要功能
## 1.可以通过该中间件,进行存储层注册( redis nosql 等. 注册的信息存储在mysql中.
## 2.可以进行删除添加存储,
## 3.代理中间会根据模块选择该模块下注册的缓存机制,
## 4.每次会对缓存服务器进行检测(这里看是否需要写个定时任务.每秒对所有的缓存做心跳检测.来判断各个存储服务器的运行情况.)

## 目标安排计划
### 1.第一步先完成数据的添加.
### 2.第二步完成数据的连接检测.
### 3.第三步实现对数据的存储.
### 4.第三步对redis的连接检测.

#### 分割线 

#### 需要重新组织下开发内容.按照以下7个模块来开发一个对外的api接口
##### 1. mysql 模块
##### 2. redis 模块
##### 3. 连接池模块
##### 4. http服务模块
##### 5. 长连接模块
##### 6. 打点日志模块
##### 7. 错误处理模块

