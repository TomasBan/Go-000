package main

import (
	"database/sql"
	"fmt"
	xerror "github.com/pkg/errors"
)

func main() {
	err := service()
	if err != nil {
		fmt.Printf("main:%+v\n",err)
	}
}


func service() error {
	// 业务代码
	return dao()
}

func dao() error {
	// db逻辑
	// ErrNoRows 说明没有匹配的数据，直接向上层暴露错误来代表没查询到数据
	return xerror.Wrapf(sql.ErrNoRows,"no rows")
}
