package main

import (
	"fmt"
	l "log"
	"os"
)

type Logger struct {
	l.Logger
}

func NewLogger() (*Logger, *os.File) {
	file, err := os.OpenFile("portfolio.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		l.Fatalf("Couldnot create a log file")
	}

	logger := &Logger{*l.New(file, "", l.Ldate|l.Ltime)}
	//logger.SetOutput(file)
	return logger, file
}

func (logger *Logger) LogInfo(data interface{}) {
	l.Println("I\t", data)
}

func (logger *Logger) LogWarning(data interface{}) {
	l.Println("W\t", data)
}

func (logger *Logger) LogError(data interface{}) {
	l.Println("E\t", data)
}
func (logger *Logger) LogInfof(format string, data ...interface{}) {
	content := fmt.Sprintf(format, data)
	l.Println("I\t", content)
}

func (logger *Logger) LogWarningf(format string, data ...interface{}) {
	content := fmt.Sprintf(format, data)
	l.Println("W\t", content)
}

func (logger *Logger) LogErrorf(format string, data ...interface{}) {
	content := fmt.Sprintf(format, data)
	l.Println("E\t", content)
}
