# 基于go-zero框架实现简单增删改查
该项目基于实际应用改造
## 创建服务
### 生成api
```shell
cd D:\go-practic\eventdispatch\app\auth\api\desc
goctl api go -api .\user.api -dir ../ -style=goZero
go mod tidy
```
### 生成model
```shell
cd D:\go-practic\eventdispatch\app\auth\model
goctl model mysql datasource -url="root:Root&20200306@tcp(IP:port)/eventdispatch" -table="sys_user"  -dir="./" -cache=true --style=goZero
```
### 生成rpc
```shell
cd D:\go-practic\eventdispatch\app\auth\rpc\pb

```