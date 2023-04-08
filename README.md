# log

## 日志库，目前提供简单的控制台输出，支持日志级别、颜色、代码位置

```go
    log := NewLogger()
    log.Info("age",12,"height","100cm")
    log.Debugf("run command err :%s",err)
```