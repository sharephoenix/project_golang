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