// @Author: Perry
// @Date  : 2020/4/26
// @Desc  : 本文件演示err code的封装

package logger

import (
	"go.uber.org/zap"
	"testing"
)

const (
	RUNERROR           = "1001"
	CFGLOADERROR       = "4001"
	INCOMPLETECFGERROR = "4002"
	UNKNOWNERR         = "9000"
)

var recodeText = map[string]string{
	RUNERROR:           "进程启动失败",
	CFGLOADERROR:       "配置文件读取失败",
	INCOMPLETECFGERROR: "配置文件项缺失",
	UNKNOWNERR:         "未知错误",
}

func TestErrCode(t *testing.T) {
	Info("a error message with code", CFGLOADERROR)
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[UNKNOWNERR]
}

/*
Output:
2020-04-26 16:50:56	INFO  zap_logger/logger_error_code_test.go:27  a error message with code {"4001": "配置文件读取失败"}
*/

/*转换errorCode到zap的field,用于Log方法封装*/
func errCodeToFields(errCodes ...string) []zap.Field {
	var fields []zap.Field
	for _, errCode := range errCodes {
		fields = append(fields, zap.String(errCode, RecodeText(errCode)))
	}
	return fields
}

/*Log方法封装*/
func Info(msg string, errCode ...string) {
	Logger.Info(msg, errCodeToFields(errCode...)...)
}
func Error(msg string, errCode ...string) {
	Logger.Error(msg, errCodeToFields(errCode...)...)
}
func Warn(msg string, errCode ...string) {
	Logger.Warn(msg, errCodeToFields(errCode...)...)
}
func Panic(msg string, errCode ...string) {
	Logger.Panic(msg, errCodeToFields(errCode...)...)
}
func Fatal(msg string, errCode ...string) {
	Logger.Fatal(msg, errCodeToFields(errCode...)...)
}
func DPanic(msg string, errCode ...string) {
	Logger.DPanic(msg, errCodeToFields(errCode...)...)
}
