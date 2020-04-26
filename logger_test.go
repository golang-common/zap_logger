// @Author: Perry
// @Date  : 2020/4/26
// @Desc  : zap_logger库的测试

package logger

import (
	"go.uber.org/zap"
	"testing"
)

/*简单的日志输出*/
func TestSimpleOutput(t *testing.T) {
	SetLevel("info") // 如果不指定日志级别,默认级别为debug
	Logger.Info("a info message")
	Logger.Debug("a debug message")
	Logger.Error("a error message with field", zap.Int("code", 400))
}

/*
Output:(因为日志级别设置为info,所以debug消息不显示)
2020-04-26 16:16:19	INFO	testing/testing.go:992	a info message
2020-04-26 16:16:19	ERROR	testing/testing.go:992	a error message with field	{"code": 400}
*/

/*初始化配置文件*/
func TestFileLog(t *testing.T) {
	ResetLogger("info", "./abc.log")
	Logger.Info("a info message")
	Logger.Debug("a debug message")
	Logger.Error("a error message with field", zap.Int("code", 400))
}
/*
StdOut Output:
2020-04-26 16:16:19	INFO	testing/testing.go:992	a info message
2020-04-26 16:16:19	ERROR	testing/testing.go:992	a error message with field	{"code": 400}

Abc.log Output:
2020-04-26 16:16:19	INFO	testing/testing.go:992	a info message
2020-04-26 16:16:19	ERROR	testing/testing.go:992	a error message with field	{"code": 400}
*/

func TestSetLevel(t *testing.T) {
	Logger.Debug("debug message 1")
	Logger.Error("error message 1")
	SetLevel("info")
	Logger.Debug("debug message 2")
	Logger.Error("error message 2")
}
/*
2020-04-26 16:37:30	DEBUG	testing/testing.go:992	debug message 1
2020-04-26 16:37:30	ERROR	testing/testing.go:992	error message 1
2020-04-26 16:37:30	ERROR	testing/testing.go:992	error message 2
*/