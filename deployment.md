# Deployment

0. 实用命令
1. Nginx
2. htop
3. Docker
4. Docker Compose
5. Docker Network
6. Images

## 0. 实用命令
```shell
# windows创建SSH密钥对
ssh-keygen -t rsa -f C:\Users\Administrator\.ssh\gcp-lz -C liuzhaomax1 -b 2048
```

## 1. Nginx
```shell
# 安装
#sudo amazon-linux-extras install nginx1
sudo yum install -y nginx

# 启动
systemctl start nginx

# 检查Active状态
systemctl status nginx

# 查看端口占用
ss -tnl
```
配置文件路径：/etc/nginx/nginx.conf

## 2. htop
资源管理工具用来查看CPU内存
```shell
# 安装
sudo yum -y install htop

# 查看
htop
```

## 3. Docker
```shell
# 查看是否已安装
yum list installed | grep docker

# 安装docker
yum -y install docker

# 启动
systemctl start docker

# 检查Active状态
systemctl status docker
```

## 4. Docker Compose
```shell
# 安装
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.26.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# 脚本文件执行权限
sudo chmod +x /usr/local/bin/docker-compose

# 查看是否安装成功
docker-compose -v

# 如果显示bash命令未找到
cp -rf /usr/local/bin/docker-compose /usr/bin/docker-compose
```

## 5. Docker Network
使用默认`bridge`

## 6. Images
pull them all.

| Service       | Images                        |
|:--------------|:------------------------------|
| DTM	          | yedf/dtm                      |
| Etcd	         | bitnami/etcd                  |
| Mysql	        | mysql:8.0                     |
| Redis	        | redis:7.0                     |
| Mysql Manage	 | phpmyadmin/phpmyadmin         |
| Redis Manage	 | erikdubbelboer/phpredisadmin  |
| Prometheus	   | bitnami/prometheus            |
| Grafana	      | grafana/grafana               |
| Jaeger	       | jaegertracing/all-in-one:1.28 |

## 7. Containers
```shell
# 根目录执行
docker-compose up -d
```

