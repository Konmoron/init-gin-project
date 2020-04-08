# init gin project

init a `gin` project.

- use `zap` as logger
- use `viper` to read config.yaml
- set healthcheck controller

## about config

some config, such as db config or password, cannot upload to github or other online website. so we need read from `OS.ENV`.

1. read config from `config.yaml`
2. read config from `OS.ENV`

the name of config in os must be SCREAMING_SNAKE_CASE. such as: the name of `config.AppName` in os is `APP_NAME`.

## References

- [go-gin 框架使用](https://juejin.im/post/5bfbbaa5e51d45315070d435)
- [一个 Gin 项目结构示例](https://www.jianshu.com/p/92919004293d)
