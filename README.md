# helpers
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gscaffold/helpers?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/gscaffold/helpers)](https://github.com/gscaffold/helpers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gscaffold/helpers)](https://goreportcard.com/report/github.com/gscaffold/helpers)
[![Unit-Tests](https://github.com/gscaffold/helpers/workflows/Go/badge.svg)](https://github.com/gscaffold/helpers/actions)
[![Coverage Status](https://coveralls.io/repos/github/gscaffold/helpers/badge.svg?branch=main)](https://coveralls.io/github/gscaffold/helpers?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/gscaffold/helpers.svg)](https://pkg.go.dev/github.com/gscaffold/helpers)

## 介绍
1. loggers: 集成常用插件, 底层基于 zap 实现.
2. app: 应用生命周期管理, 包括启动、配置、初始化、销毁.
3. rest: http 相关功能封装. 支持通过 app 管理 http 服务.
4. rpc: rpc 相关服务封装. 支持通过 app 管理 rpc 服务, 支持 grpc server/client.

组件的使用事例参考 [examples](./examples)

## Install
```shell
go get github.com/gscaffold/helpers
```
