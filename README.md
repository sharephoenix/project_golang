# kratos-demo 是主要的工程

[README](./kratos-demo/README.md)
[PROBLEM](./Problem.md)

## go 平滑启动服务
* go 自带平滑重启功能
* k8s 部署，实现平滑重启
* jenkins 实现构建

## gin
* 完成

## gin + jwt 基本实现
* 完成

## gin 中间件 时间 api 基本校验，middleWareHandle
* 完成

## 微服务化
* mac 安装 **docker edge** 版本（自带 k8s）
* 自动发布 docker 到 hub.docker.com 仓库
* 重新应用 service eg: ```kubectl apply -f xxx.yaml```
* 先完成业务再研究 centos 的 k8s 安装 和 jenkin 的自动交付功能