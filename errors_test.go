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
	err := CallErrTypeFn("0") // 測試套件的error
	if err != nil {
		fmt.Println("返回給前端的錯誤 code :", codeMessage.GetErrorCode(err, true))
	}
	fmt.Println()

	err = CallErrTypeFn("1") // 測試原生的error
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

func CallErrTypeFn(id string) error {
	if id != "" {
		return ErrDataInvalid.Write("錯誤內容: 'id' 不能是預設值" + id)
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
