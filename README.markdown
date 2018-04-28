# Cobweb

> 分布式爬虫， 目前支持珍爱网的爬取

### 特点

- 分布式
- 使用ElasticSearch作为存储
- 自带GUI洁面


### 开始使用

安装依赖:
```shell
go get github.com/urfave/cli

go get golang.org/x/text
go get golang.org/x/net

go get gopkg.in/olivere/elastic.v5
```

编译:
```shell
go build
```

启动存储服务:
`cobweb saver --port=8700`

启动worker服务组:
```shell
cobweb worker --port=7900
cobweb worker --port=7901
cobweb worker --port=7902
```

启动调度器:
`cobweb dispatch --saver-port=8700 --worker-ports=7900,7901`

启动Web GUI:
`cobweb web --port=8091`

### Todo:
 
 - 去重持久化
 - IP Proxy
 - User Agent Pool
 - 配置化
 - 抽象parser接口，便于接入更多网站的爬虫业务


