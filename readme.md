## goship 框架模板
golang 开发 api接口 基本框架

## 依赖
- golang   
- swagger  

### golang 安装
golang 官网: https://golang.org/  
如果没有开启 go module,则需要开启,可以通过 `go env` 查看`GO111MODULE`开启情况  

开启命令
```shell script
export GO111MODULE=on
```

### 安装 swagger
`gin-swagger`官网: https://github.com/swaggo/gin-swagger  

简洁安装
```shell script
go get -u github.com/swaggo/swag/cmd/swag
```
若 `$GOPATH/bin` 没有加入$PATH中，你需要执行将其可执行文件移动到$GOBIN下
```shell script
mv $GOPATH/bin/swag /usr/local/go/bin
```

## 运行
```bash
# 切换到项目根目录
cd ~/path/to/gopro

# 安装更新依赖
go mod tidy && go mod download

# 生成 api 文档
swag init

# 修改配置,只需要修改 config.toml 配置即可
cp config.toml.example config.toml

# 运行
go run main.go -f config.toml
```
> 配置文件,重点修改数据库配置和http端口号 
