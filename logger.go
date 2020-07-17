package main

import (
	l "log"
	"os"
	"os/exec"
)

type Logger struct {
	InfoLogger    *l.Logger
	WarningLogger *l.Logger
	ErrorLogger   *l.Logger
}

func NewLogger() (*Logger, *os.File) {
	logger := &Logger{}
	f, err := os.OpenFile("/var/log/portfolio.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		cmd := exec.Command("sudo", os.Args...)
		b, err := cmd.CombinedOutput()
		if err != nil {
			l.Fatalf("Could not create log file. Error %v", err)
		}
		logger.InfoLogger.Printf("%s", b)
	}
	//	defer f.Close()
	logger.InfoLogger = l.New(f, "I\t", l.LstdFlags|l.Lshortfile)
	logger.WarningLogger = l.New(f, "W\t", l.LstdFlags|l.Lshortfile)
	logger.ErrorLogger = l.New(f, "E\t", l.LstdFlags|l.Lshortfile)

	return logger, f
}

func (logger *Logger) LogInfo(data interface{}) {
	logger.InfoLogger.Println(data)
}
func (logger *Logger) LogInfof(format string, data ...interface{}) {
	logger.InfoLogger.Printf(format, data)
}

func (logger *Logger) LogWarning(data interface{}) {
	logger.WarningLogger.Println(data)
}
func (logger *Logger) LogWarningf(format string, data ...interface{}) {
	logger.WarningLogger.Printf(format, data)
}

func (logger *Logger) LogError(data interface{}) {
	logger.ErrorLogger.Println(data)
}
func (logger *Logger) LogErrorf(format string, data ...interface{}) {
	logger.ErrorLogger.Printf(format, data)
}
func (logger *Logger) LogFatal(data interface{}) {
	logger.ErrorLogger.Fatal(data)
}
