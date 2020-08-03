# 初始化 `gin` 项目

## 文档

- [english](./readme.md)
- [中文](./readme_zh.md)

## 简介

初始化一个 `gin` 项目。

- 使用 `zap` 输出日志
- 使用 `viper` 读取 `config.yaml`
- 设置 `healthcheck` 路由
- 为 `healthcheck` 路由增加单元测试

## 关于配置文件

一些配置选项，比如数据库信息、密码等，不能上传到公开的网络环境，因此可以从系统环境变量读取。

配置的加载顺序为：

1. 从 `config.yaml` 读取
2. 从 系统环境变量 中再次读取

`config` 中的配置项的名称，在系统环境变量的对应名称格式为：SCREAMING_SNAKE_CASE，比如：`config.AppName`，在系统的环境变量名称为：`APP_NAME`。

## 关于单元测试

```shell
cd test
go test -v -run TestHealthcheckRoute
```

## about swagger

- [swagger usage](https://github.com/swaggo/swag)
- [gin-swagger usage](https://github.com/swaggo/gin-swagger)

## 参考

- [go-gin 框架使用](https://juejin.im/post/5bfbbaa5e51d45315070d435)
- [一个 Gin 项目结构示例](https://www.jianshu.com/p/92919004293d)
