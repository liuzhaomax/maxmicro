# AWS

1. Nginx
2. htop
3. Docker
4. Docker Compose

## 1. Nginx
```
//安装
sudo amazon-linux-extras install nginx1

//启动
systemctl start nginx

//检查Active状态
systemctl status nginx

//查看端口占用
ss -tnl
```

## 2. htop
资源管理工具用来查看CPU内存
```
//安装
yum -y install htop

//查看
htop
```

## 3. Docker

```
//查看是否已安装
yum list installed | grep docker

//安装docker
yum -y install docker

//启动
systemctl start docker

//检查Active状态
systemctl status docker
```

## 4. Docker Compose

```
//安装
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

//脚本文件执行权限
sudo chmod +x /usr/local/bin/docker-compose

//查看是否安装成功
docker-compose -v
```
