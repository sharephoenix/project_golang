#!/bin/bash
#mac 构建方法
#go build -ldflags "-w" user.go

echo "========== 开始构建 =========="
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build user.go
echo "========== 二进制信息 =========="
ls -al user
echo "========== 开始打包 docker =========="
docker build -t luanys/zcdocker:1.0.1 .
echo "========== 开始推送 docker 到远端 =========="
docker push luanys/zcdocker:1.0.1
echo "========== end =========="