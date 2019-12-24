# Docker 

## nginx docker 相关命令
[相关命令](https://www.jianshu.com/p/9fd2f77001a3)
* 拉取 image
docker pull nginx
docker pull redis
* 查看当前运行的 docker
```
docker ps
docker ps -a
```
* 当前安装的images
```
docker images
```

* 查看当前 continerid 的 唯一值
1. practical_rosalind 为运行container 的 name；使用 docker ps 查看
```
docker inspect -f '{{.ID}}' practical_rosalind
```
* 把当前环境中的文件 copy 到 docker 中
```
docker cp /Users/apple/ClientsRemote/xp-telephone-book/dist/ 043b601c8aa29190e452880fee6c8192415c2bba9e43ba1934ab99c01435dfa5:/usr/share/nginx/html/
```
* 进入某个容器的控制台
1. 043b601c8aa2 当前容器id
```
docker exec -it 043b601c8aa2 /bin/bash
```
* 运行一个 docker 容器
1. 80 为容器内部的 端口号；8000 为当前服务器的端口号
2. -d 表明在后端运行
```
docker run -p 8000:80 -d nginx
```

* 停止 docker 容器
```
docker stop 043b601c8aa2
```

* 移除 docker 容器
```
docker rm 043b601c8aa2
docker containId rm 043b601c8aa2
```

* 启动挂载宿主机文件夹
1. -v 宿主机文件:镜像文件目录
```
docker run -p 8000:80 -v /Users/apple/Downloads/html:/usr/share/nginx/html -d nginx
```

## Docker-Compose 集群
[友情链接](https://blog.51cto.com/9291927/2310444)

## Docker 部署 golang
[友情链接](https://www.jianshu.com/p/5939dcf5c96e)

## 启动 mysql
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=qwe123 -d mysql
