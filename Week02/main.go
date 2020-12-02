package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	if err := Dao(0); err != nil {
		err = errors.Wrap(err, "查询错误")
		fmt.Printf("%+v\n", err)
	}
	fmt.Println("exit...")
}

// Dao 数据库操作
func Dao(id int64) (err error) {
	//假设这里查询数据库，然后返回错误
	err = sql.ErrNoRows
	if nil != err {
		err = errors.Wrap(err, "数据未找到")
		return
	}
	return
}
