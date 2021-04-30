# 处理依赖注入
https://juejin.cn/post/6844903901469097998
* 需要go get github.com/google/wire/cmd/wire
* 只需要构建wire.go ，通过执行命令：wire 生成 wire_gen.go
* 执行go run main.go wire_gen.go 才能跑起来，因为main.go 引入了wire_gen.go 的方法
