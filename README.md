###我也不知道是第几天了，对micro又有了重新的认识，相比较springcloud而言，micro更灵活，更健全###
###micro.newservice 一个服务之后，相当于java写一个service，但是这个service是没法调用的，这个没法调用是指没有暴露出去，外部没有办法通过url去调用###
###所以代码里用的是gin，go的一个web框架，利用route路由，也就是生成url 去调用写的service，还有一种是用micro工具 生成网关，对外提供一个统一的地址暴露出去###
###设置参数，有点类似springcloud，的服务名称###
set MICRO_REGISTRY=etcd
set MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
set MICRO_API_NAMESPACE=yuxiang.li
set MICRO_API_HANDLER=rpc
micro api

###整合数据库gorm###