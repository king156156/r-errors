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

	// 非本套件的error處理 ex: 原生error
	mylog.Error(err.Error())
	return
}
