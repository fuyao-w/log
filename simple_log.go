package log

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	. "github.com/fuyao-w/common-util"
)

var (
	ResetColor string
	Red        string
	Green      string
	Yellow     string
)
var (
	gopath = path.Join(os.Getenv("GOPATH"), "src") + "/"
)

type Level int

const (
	Debug Level = iota + 1
	Waring
	Info
	Error
)

func (l *Level) String() (str string) {
	switch *l {
	case Debug:
		str = "Debug"
	case Waring:
		str = "Waring"
	case Info:
		str = "Info"
	case Error:
		str = "Error"
	default:
		str = "Unknown"
	}
	return
}

var logColor = map[Level]string{
	Debug:  Green,
	Waring: Yellow,
	Info:   Green,
	Error:  Red,
}

func buildPrefix(level Level) *bytes.Buffer {
	toBuf := Str2Bytes
	buf := new(bytes.Buffer)
	buf.Write(toBuf(logColor[level]))
	buf.Write(toBuf(level.String()))
	buf.Write(toBuf("|"))
	buf.Write(toBuf(printCodeLocation()))
	buf.Write(toBuf("| "))
	buf.Write(toBuf(ResetColor))
	return buf
}

type simpleLog struct {
	log *log.Logger
}

func NewLogger() *simpleLog {
	return &simpleLog{log: log.Default()}
}
func printCodeLocation() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return "unknown location"
	}
	if strings.HasPrefix(file, gopath) {
		file = file[len(gopath):]
	}
	return fmt.Sprintf("%s.%d", file, line)
}
func (s *simpleLog) Infof(format string, v ...any) {
	s.log.Printf(buildPrefix(Info).String()+format, v...)
}

func (s *simpleLog) Info(v ...any) {
	s.log.Println(append([]any{any(buildPrefix(Info).String())}, v...)...)
}

func (s *simpleLog) Errorf(format string, v ...any) {
	s.log.Printf(buildPrefix(Error).String()+format, v...)
}

func (s *simpleLog) Error(v ...any) {
	s.log.Println(append([]any{any(buildPrefix(Error).String())}, v...)...)
}

func (s *simpleLog) Warnf(format string, v ...any) {
	s.log.Printf(buildPrefix(Waring).String()+format, v...)
}

func (s *simpleLog) Warn(v ...any) {
	s.log.Println(append([]any{any(buildPrefix(Waring).String())}, v...)...)
}

func (s *simpleLog) Debugf(format string, v ...any) {
	s.log.Printf(buildPrefix(Debug).String()+format, v...)
}

func (s *simpleLog) Debug(v ...any) {
	s.log.Println(append([]any{any(buildPrefix(Debug).String())}, v...)...)
}
