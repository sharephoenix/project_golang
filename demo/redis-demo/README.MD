# redis 使用

## 安装
[README](https://www.jianshu.com/p/bb7c19c5fc47)

## 配置
[redis.conf](https://www.cnblogs.com/zxtceq/p/7676911.html)
1. 配置文件 redis.conf
2. save 10 1 -- 10s 中保存一次

## golang redis 操作

[学习链接](https://www.cnblogs.com/wdliu/p/9330278.html)

## redis 操作命令

* 启动 redis 服务
```
redis-server 
```

* 启动 redis-cli
```
redis-cli
```

* 当前所有的 keys
```
keys *
```

* 存储数据

```
set key0 value0
```

* 获取存储值
```
get key0
```

## 三方

```
demo 使用 github.com/go-redis/redis 连接 set get 数据
```

## uuid
```
"github.com/satori/go.uuid"
```

## 服务开启只需要一个 redis