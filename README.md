# go-gin-base

## go gin 基础框架

### 库

- https://github.com/go-validator/validator 验证参数
- https://github.com/go-playground/validator 验证参数
- github.com/robfig/cron 消息队列
- github.com/sirupsen/logrus log
- github.com/streadway/amqp mq

### Go Module代理

~~~
GOPROXY=https://proxy.golang.org    # 官方推荐，国内还无法正常使用
GOPROXY=https://goproxy.cn          # 国内相对友好
GOPROXY=https://goproxy.io          # 通用
~~~