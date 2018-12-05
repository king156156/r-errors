package errors_test

import (
	"fmt"
	"log"
	"testing"

	errors "github.com/king156156/r-errors"
)

type MyIErr interface {
	Error() string
}

type MyErr struct {
	str string
}

func NewMyErr(msg string) MyIErr {
	return &MyErr{msg}
}

func (my *MyErr) Error() string {
	return my.str + " err, err"
}

var s = errors.MsgCode{
	fnerr: "11110",
}

var fnerr = NewMyErr("Error Test MyErr")

func Test_SetErrorfn(t *testing.T) {
	// 自訂方法,調用方式, 可塞自己想要觸發的log套件
	errors.SetErrorfn(func(err error) bool {
		if me, ok := err.(MyIErr); ok {
			log.Println("自訂log:", me.Error())
			return true
		}
		return false
	})

	// 開始測試
	err := getErr()

	// 返回前端error code
	errCode := s.GetErrorCode(err, true)
	fmt.Println("吐回前端  errCode:", errCode)
}

func getErr() error {
	return fnerr
}
