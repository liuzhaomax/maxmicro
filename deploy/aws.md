# AWS

1. Nginx
2. htop

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
yum install htop

//查看
htop
```