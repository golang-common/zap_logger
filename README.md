# zap日志库封装
本工具针对zap日志库的封装

## 安装
`go get github.com/lyonsdpy/zap_logger`

## 使用
简单的使用方式如下,其中
```
import logger "github.com/lyonsdpy/zap_logger"

func main(){
    logger.Info("a log message")
    logger.Info("a log message with field")
}
```
初始化配置文件配置,如下代码会在当前目录中生成abc.log文件并重新设置log level.
```
import logger "github.com/lyonsdpy/zap_logger"

func main(){
    ResetLogger("info", "./abc.log")
    Logger.Info("a info message")
    Logger.Debug("a debug message")
    Logger.Error("a error message with field", zap.Int("code", 400))
}
```
可以随时调用`SetLevel`来动态调整日志级别
本封装已经包含了log rotate,默认参数如下
- 文件最大多少M,默认100M
- 最多保留多少天,默认不根据日期删除,默认7天
- 最多保留多少个备份,默认7个
- 是否使用本地时间,默认使用UTC时间
- 是否gzip压缩

