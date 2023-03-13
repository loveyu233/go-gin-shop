# 项目介绍
## 目录展示
````
|-- api
|-- core
|-- enter
|-- global
|-- initialize
|-- middleware
|-- redisLock
|-- resources
|-- router
|-- servic
|-- static
|   |-- casbin
|   |-- code
|-- utils
````
## 项目实用技术:
>Go,Gin,Mysql,Redis,MongoDB

## 项目主要实现功能
1. 使用redis作缓存
2. 解决redis缓存产生的缓存雪崩缓存击穿和缓存穿透问题
3. 实现优惠卷的秒杀
4. 使用lua脚本+redis实现分布式锁解决超卖问题
5. 使用Redis实现了用户点赞,关注,共同关注等功能
6. 使用MongoDB实现评论和回复功能