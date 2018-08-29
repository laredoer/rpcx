# rpcx

wule61/rpcx 是基于 go 语言rpc开源项目的实践

rpcx 的优势

1.  易学习，适合快速开发

    > 易于入门, 易于开发, 易于集成, 易于发布,易于监控，不用写protoc协议

2.  高性能

    > 性能远远高于 Dubbo、Motan、Thrift等框架，是gRPC性能的两倍

3.  交叉平台，交叉语言

    > 可以容易部署在Windows/Linux/MacOS等平台，支持各种编程语言的调用，利用docker可以方便部署

4.  服务发现

    > 除了本地外，还支持 Zookeeper、Etcd、 Consul、mDNS等注册中心

5.  服务治理

    > 支持 Failover、 Failfast、 Failtry、Backup等失败模式，支持 随机、 轮询、权重、网络质量, 一致性哈希,地理位置等路由算法

6.  丰富功能

    > 配合gin使用，可扩展很多功能

