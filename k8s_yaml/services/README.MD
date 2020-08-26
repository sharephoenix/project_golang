[中文文档](https://www.kubernetes.org.cn/kubernetes-services)
## 常用命令
* kubectl get svc -n new-namespace
* kubectl get pods
* kubectl get namespace
* kubectl log podName

## docker 常用命令
* docker build -t luanys/services-user:1.0.0 .
* docker run -p port:containerPort --name containerName -d imageId
* docker rm containerId
* docker rmi imageId
* docker stop containerId
* docker start containerId


## DockerFile
```ruby
FROM nginx

MAINTAINER lanys "5507098xx@qq.com"

WORKDIR /src/godocker

ADD api /src/godocker

EXPOSE 9090
ENTRYPOINT ["./user", "-f", "./etc/config.json"]
```

## 发布 docker image

```shell
#!/bin/bash
#mac 构建方法
#go build -ldflags "-w" user.go

echo "========== 开始构建 =========="
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api/user ./api/user.go
echo "========== 二进制信息 =========="
ls -al api/user
echo "========== 开始打包 docker =========="
docker build -t luanys/services-user:1.0.0 .
echo "========== 开始推送 docker 到远端 =========="
docker push luanys/services-user:1.0.0
echo "========== end =========="
```

## 总结
* dockerImage 不修改版本号 `kubectl delete -f xxx` 后， 再执行 `kubectl apply -f xxx` 会自动更新镜像
* dockerImage 修改版本号后， **xxx.yaml** 再执行 `kubectl apply -f xxx.yaml` 也会自动更新
