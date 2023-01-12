## 简介

gin写的一个博客

## 环境的安装

### 安装扩展

```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/spf13/viper
go get -u gopkg.in/natefinch/lumberjack.v2
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/go-playground/validator/v10
go get -u github.com/golang-jwt/jwt/v4
go get -u gopkg.in/gomail.v2
go get -u github.com/juju/ratelimit
```

### 安装swag命令行工具

```
go install github.com/swaggo/swag/cmd/swag@latest
```

### swag生成文件

在项目目录下执行如下命令

```
swag init
```

### 安装jaeger

```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -e COLLECTOR_OTLP_ENABLED=true \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.41
```

### 访问jaeger

```
http://localhost:16686/
```

### 由于这个 github.com/uber/jaeger-client-go 库停止维护了

### 建议使用 github.com/open-telemetry/opentelemetry-go 新库

```
go get -u go.opentelemetry.io/otel
go get -u go.opentelemetry.io/otel/exporters/jaeger
```

### sql追踪

```
go get -u gorm.io/plugin/opentelemetry
go get -u gorm.io/plugin/opentelemetry/metrics
```

### 把配置文件打包进二进制文件中

```
go get -u github.com/go-bindata/go-bindata/...
go install github.com/go-bindata/go-bindata/...
```

生成配置文件go代码，-pkg表示生成的package name为configs

```
go-bindata.exe -o configs/config.go -pkg=configs configs/config.yaml
```

通过如下代码就可以获取配置文件内容

```go
asset, _ := configs.Asset("configs/config.yaml")
fmt.Println(string(asset))
```

### 配置热更新

```
go get -u github.com/fsnotify/fsnotify
```