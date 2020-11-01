# docker-mon 
作为Nightingale的插件，用于收集docker容器的监控指标

## 快速构建 
```
    $ mkdir -p $GOPATH/src/github.com/n9e
    $ cd $GOPATH/src/github.com/n9e
    $ git clone https://github.com/n9e/docker-mon.git
    $ cd docker-mon
    $ ./build
    $ ./docker-mon
```

## 前置依赖
1. docker容器所在宿主机已安装并启动了cadvisor
2. docker容器的环境变量中包含 N9E_NID ，N9E_NID 的内容为夜莺服务树节点id

## 使用方式
将 docker-mon、docker-mon.yml 分发到容器所在的宿主机上，到宿主机所在的节点配置插件采集即可
![docker-mon](https://s3-gz01.didistatic.com/n9e-pub/image/docker.png)