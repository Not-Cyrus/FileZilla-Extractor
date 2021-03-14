package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Debug(v error) {
	if formatLogger.LogLevel == "error" {
		Error(v)
	}
	fmt.Println(v)
}

func Error(v error) {
	panic(v)
}

func HandleError(err error) {
	if err != nil {
		Debug(err)
	}
}

func InitLog(level string) {
	formatLogger = newLog(os.Stdout).setLevel(level)
}

func newLog(w io.Writer) *Logger {
	return &Logger{
		log: log.New(w, "", 0),
	}
}

func (l *Logger) setLevel(level string) *Logger {
	l.LogLevel = level
	return l
}

type (
	Logger struct {
		LogLevel string
		log      *log.Logger
	}
)

var (
	formatLogger *Logger
)
