
<img align="right" width="380"  src="./docs/lemon.jpeg">

# lemon

English | [简体中文](README.md)

[![Go](https://img.shields.io/badge/go-1.16.3-green)](https://github.com/golang/go)
[![lemon](https://img.shields.io/badge/lemon-v0.2-blue)](http://subway_intro)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](http://subway_intro/blob/main/LICENSE)


## 简介

[lemon](http://subway_intro) 是一个hdf自研的Golang DDD框架。
支持[Grpc](https://github.com/grpc/grpc)，[Grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)集成了注册中心[Eureka](https://github.com/Netflix/eureka),集成了丰富的Grpc中间件，[sentry](https://github.com/getsentry/sentry),[Zap](https://github.com/uber-go/zap),[Jaeger](https://github.com/jaegertracing/jaeger)

## 前序准备

你需要在本地安装 [go](https://github.com/golang/go) 

## Quick Start

```bash
# Install from source
cd $GOPATH/src && git clone http://subway_intro.git && cd lemon

# 初始化
go mod tidy

# 本地执行
go run main.go server

```
**help 信息**
```
go run main.go help

```
---

## Change log
**v0.1 2021-02-01**
 - DDD 目录结构
 - 支持EventBus
 - 最小Domain

---

**v0.2 2021-08-11**
- 【修复】 logger level 在特殊情况下失效的bug

- 【优化】 server 启动模式，支持平滑退出
- 【优化】 eventbus 支持异步订阅

- 【新增】 grpc
- 【新增】 ttp grpc-gateway
- 【新增】 grpc中间件支持tracer
- 【新增】 grpc中间件支持logger
- 【新增】 grpc中间件支持timeout
- 【新增】 grpc中间件支持sentry
- 【新增】 grpc中间件支持recovery
- 【新增】 eureka注册中心
- 【新增】 sentry
- 【新增】 logger提供zap.Logger模式

- 【重构】 internal.app 统一中间件注册模式 注入到context中
- 【重构】 orm模块 支持context模式
- 【重构】 eventbus 支持context模式
- 【重构】 eureka 支持context模式
- 【重构】 logger 支持context模式

- 【移除】 gin模块
- 【移除】 sidecar依赖

---

## 目录结构说明
```
├── Dockerfile
├── LICENSE
├── Makefile  # CI/CD
├── README.md
├── app # 程序入口
│   ├── daemon
│   │   └── cobra.go # 入口文件
│   ├── mq  # mq 支持rabbitmq / kafka
│   │   ├── init.go
│   │   └── rabbitmq
│   │       └── build.go
│   ├── server  # grpc / grpc-gateway 入口
│   │   └── init.go
│   └── util # response for json
│       └── util.go
├── build  # 二进制 可执行文件
├── internal # 核心领域
│   ├── app.go # 核心领域 context
│   ├── client # grpc 调用第三方服务
│   │   └── client.go
│   ├── constant # 扩展的const 常量
│   │   └── greeter.go
│   ├── domain # 领域层
│   │   ├── greeter # 按聚合根组织文件夹
│   │   │   └── greeter.go # greeter entity
│   │   ├── domain.go # 通用抽象接口 定义聚合根 和 值对象 行为
│   │   └── errors.go # 核心领域 错误码
│   ├── dto # 领域事件中扩展的content 可用于聚合多实体字段发送领域事件
│   │   └── dto.go # greeter dto
│   ├── event # 领域内事件总线 支持同步或异步
│   │   ├── greeter_event_handle.go # 领域事件订阅者greeter
│   │   └── register_handle.go # 领域事件订阅注册器
│   ├── infrastructure # 基础支持 一般只用于当前项目的工具集或中间件，等稳定了或者通用了可抽离到中间件层
│   │   ├── domain # 处理领域通用逻辑
│   │   │   ├── error.go # 处理错误的handler
│   │   │   └── pagination.go # 处理分页
│   │   └── tools # 当前项目工具箱 或不稳定的中间件
│   │       ├── config # 配置处理中心
│   │       │   ├── application.go # 当前项目主配置
│   │       │   ├── clickhouse.go # hdf-clickhouse 配置
│   │       │   ├── client.go # hdf-grpc 调用第三方配置
│   │       │   ├── config.go # 配置处理核心类
│   │       │   ├── database.go # hdf-orm 驱动配置
│   │       │   ├── eureka.go # 注册中心配置
│   │       │   ├── evbus.go # 事件订阅者配置
│   │       │   ├── jwt.go # hdf jwt 权限配置
│   │       │   ├── kafka.go # hdf kafka 配置
│   │       │   ├── logger.go # hdf-zap.logger 配置
│   │       │   ├── rabbitmq.go # hdf-rabbitmq 配置
│   │       │   ├── rpc.go # hdf-grpc 配置
│   │       │   ├── sentry.go # hdf-sentry 配置
│   │       │   └── tracer.go $ hdf-jaeger 配置
│   │       ├── env.go # 环境变量
│   │       └── string.go # 常用字符串处理函数
│   ├── interfaces # 核心领域接口定义
│   │   └── greeter.go # greeter interface
│   ├── repos # 仓储
│   │   ├── greeterrepo.go # 仓储greeter
│   │   └── factory.go # 仓储初始化工厂
│   ├── service # 业务逻辑层
│   │   ├── greeter.go # 业务greeter
│   │   ├── greeter_test.go # 单测
│   │   ├── service.go # 业务逻辑入口
│   │   └── service_test.go # 单测
│   └── version.go # 当前项目版本
├── docker # 用于k8s部署
│   ├── deployment.yaml
│   ├── go-fluentbit-cm-prod.yaml
│   ├── go-fluentbit-cm-test.yaml
│   └── server.yaml
├── docker-compose.yml # 用于本地docker部署
├── docs # 当前项目docs swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod # 当前项目依赖包
├── go.sum # 当前项目依赖包
├── main.go # 当前项目入口
├── proto # 当前项目 grpc and grpc-gateway 定义
│   ├── lemon.proto # 当前项目 grpc and grpc-gateway 定义
│   └── google # google http 扩展 for grpc-gateway
│       └── api
│           ├── annotations.proto
│           ├── http.proto
│           └── httpbody.proto
├── settings.dev.yml # 开发环境配置
├── settings.test.yml # 测试环境配置
└── stub # grpc 调用第三方的桩代码
    └── lemon # grpc greeter
        ├── lemon.pb.go
        ├── lemon.pb.gw.go
        └── lemon_grpc.pb.go
```

## License

[MIT](http://subway_intro/blob/main/LICENSE)

Copyright (c) 2021-present hdf

## Roadmap
[Roadmap](./roadmap.md)

