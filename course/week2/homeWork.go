package main

/*
用pkg/errors包，将调用dao层返回的sql.ErrNoRows进行wrap（可以封装一些提示信息），不做其他处理直接将error抛给上层
*/


import (
	"fmt"
	"github.com/pkg/errors"
)

//import "errors"

func queryDb(sql string) error {
	//return errors.New("sql.ErrNoRows")
	return errors.New("sql.ErrNoRows")
}

func queryDbApi() error {
	sql := "select * from table where name='test'"
	err := queryDb(sql)
	if err != nil {
        return errors.WithMessage(err, sql)
        //return errors.Wrap(err, sql)
	}
	return nil
}

func main() {
	err := queryDbApi()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
	}
}