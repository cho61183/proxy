# proxy
数据中间件的代理.

# 主要功能
## 1.可以通过该中间件,进行存储层注册( redis nosql 等. 注册的信息存储在mysql中.
## 2.可以进行删除添加存储,
## 3.代理中间会根据模块选择该模块下注册的缓存机制,
## 4.每次会对缓存服务器进行检测(这里看是否需要写个定时任务.每秒对所有的缓存做心跳检测.来判断各个存储服务器的运行情况.)

