package errors

import (
	"os"
)

const (
	// 起始的 caller 層數
	startCallerLevel = 2
)

var mylog = &Log{
	writer: os.Stdout,
}

// 自訂義格式
var inErrorfn func(error) bool

// MsgCode 錯誤訊息定義區
type MsgCode map[error]string

// GetErrorCode 判斷錯誤訊息取得 error code
func (m *MsgCode) GetErrorCode(err error, islog bool) (errorCode string) {
	// 判斷是否有註冊 error code
	errorCode, _ = (*m)[err]
	if !islog {
		return
	}

	e, ok := err.(RError)
	if ok {
		mylog.Error(e.Err().ErrorPath())
		e.Clear()
		return
	}

	if inErrorfn != nil && inErrorfn(err) {
		return
	}

	// 非本套件的error處理 ex: 原生error
	mylog.Error(err.Error())
	return
}

// SetErrorfn 設定自訂義Error處理
func SetErrorfn(fn func(error) bool) {
	inErrorfn = fn
}
