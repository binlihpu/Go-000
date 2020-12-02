/*
Week02 作业题目：
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/


个人见解：
数据库操作的错误返回最好Wrap后再返回上层，这样做的好处是上层结构在判断错误是不用引入新的数据包`"database/sql"`，这样依赖更“清真”些。另外就是Wrap以后，从代码中我可以看到，在err后我可以从打印的错误日志中看出错误的原始位置，以及后续的堆栈信息，一层一层的错误往上传，这样虽然错误的信息多点（毕竟报错的地方还是很少的），但是整个日志链路很清晰，便于定位问题。
```go
$ go run main.go
sql: no rows in result set
数据未找到
main.Dao
        /Users/Leon/work_project/golang/Go-000/Week02/main.go:23
main.main
        /Users/Leon/work_project/golang/Go-000/Week02/main.go:11
runtime.main
        /Users/Leon/.g/go/src/runtime/proc.go:200
runtime.goexit
        /Users/Leon/.g/go/src/runtime/asm_amd64.s:1337
查询错误
main.main
        /Users/Leon/work_project/golang/Go-000/Week02/main.go:12
runtime.main
        /Users/Leon/.g/go/src/runtime/proc.go:200
runtime.goexit
        /Users/Leon/.g/go/src/runtime/asm_amd64.s:1337
exit...
```

