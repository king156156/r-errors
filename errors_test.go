package errors_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	derrors "github.com/king156156/r-errors"
)

var codeMessage = derrors.MsgCode{
	ErrDataInvalid: "10001",
	ErrConnect:     "10002", // 可註冊原生errors.new
	ErrUserInvalid: "10003",
}

func Test_Code(t *testing.T) {
	err := CallErrTypeFn(0) // 測試套件的error
	if err != nil {
		fmt.Println("返回給前端的錯誤 code :", codeMessage.GetErrorCode(err, true))
	}
	fmt.Println()

	err = CallErrTypeFn(1) // 測試原生的error
	if err != nil {
		fmt.Println("返回給前端的錯誤 code :", codeMessage.GetErrorCode(err, true))
	}
	fmt.Println()

	err = A(-999) // 測試func call func的錯誤
	if err != nil {
		fmt.Println("返回給前端的錯誤 code :", codeMessage.GetErrorCode(err, true))
	}
}

// # 自訂的錯誤
var (
	ErrDataInvalid = derrors.New("Error! Data Invalid")
	ErrConnect     = errors.New("Error! 原生err")
	ErrUserInvalid = derrors.New("Error! User Invalid")
)

// # 測試1

func CallErrTypeFn(id int) error {
	if id == 0 {
		return ErrDataInvalid.Write("錯誤內容: 'id' 不能是預設值" + string(id))
	}
	return ErrConnect
}

// # 測試2 規則: 自己判斷錯誤,可以加入Write函數來寫需要顯示的參數,如果是調用其他方法在同樣規則下則不要在紀錄,不然訊息會很亂
func A(id int) error {
	if id == 0 {
		return ErrDataInvalid.Write("'id' 不能是預設值" + string(id))
	}

	if err := B(id); err != nil {
		return err
	}

	return nil
}

func B(id int) error {
	if id == -999 {
		s := strconv.FormatInt(int64(id), 10)
		return ErrUserInvalid.Write("無效的使用者 id:" + s)
	}
	return nil
}

// # 測試數據

// ➜  r-errors git:(master) ✗ go test -run="Test_Code"
// 2018/11/27 17:56:35 logerr:
// Error! Data Invalid 錯誤內容: 'id' 不能是預設值
// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:48

// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:19

// /usr/local/go/src/testing/testing.go:777

// /usr/local/go/src/runtime/asm_amd64.s:2361

// /Users/robert/go/src/github.com/king156156/r-errors/code.go:29

// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:21

// /usr/local/go/src/testing/testing.go:777

// /usr/local/go/src/runtime/asm_amd64.s:2361

// 返回給前端的錯誤 code : 10001

// 2018/11/27 17:56:35 logerr:
// Error! 原生err
// 返回給前端的錯誤 code : 10002

// 2018/11/27 17:56:35 logerr:
// Error! User Invalid 無效的使用者 id:-999
// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:69

// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:59

// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:31

// /usr/local/go/src/testing/testing.go:777

// /usr/local/go/src/runtime/asm_amd64.s:2361

// /Users/robert/go/src/github.com/king156156/r-errors/code.go:29

// /Users/robert/go/src/github.com/king156156/r-errors/errors_test.go:33

// /usr/local/go/src/testing/testing.go:777

// /usr/local/go/src/runtime/asm_amd64.s:2361

// 返回給前端的錯誤 code : 10003
// PASS
// ok      github.com/king156156/r-errors  0.006s
