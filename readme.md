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
2. docker容器的环境变量中包含 N9E_NID ，N9E_NID 的内容为夜莺服务树节点id，如果设置 N9E_NID = 1，则到节点id为1的节点下，就可以容器的监控指标

## 使用方式
1. 将 docker-mon、docker-mon.yml 分发到容器所在的宿主机上
2. 到宿主机所属节点配置插件采集

![docker-mon](https://s3-gz01.didistatic.com/n9e-pub/image/docker.png)

3. 配置完之后，到即时看图，选择对应的节点，再选择设备无关，即可查看采集到的容器监控指标
![docker-metric](https://s3-gz01.didistatic.com/n9e-pub/image/docker_metric.png)

## 视频教程
[观看地址](https://s3-gz01.didistatic.com/n9e-pub/video/n9e-docker-mon.mp4)