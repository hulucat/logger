package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
}

var logFile, _ = openLogFile()
var logger *log.Logger
var logDay int

func openLogFile() (f *os.File, err error) {
	t := time.Now()
	day := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	filename := fmt.Sprintf("./app_%s.log", day)
	logDay = t.Day()
	f, err = os.OpenFile(filename, os.O_APPEND, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(filename)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				f.Close()
				f, err = os.OpenFile(filename, 766, os.ModeAppend)
			}
		}
	}
	return
}

func getLogger() (l *log.Logger, err error) {
	if time.Now().Day() != logDay {
		logFile.Close()
		logFile, err = openLogFile()
		if err != nil {
			fmt.Errorf("Error open log file: %s\n", err.Error())
			return
		}
		l = log.New(logFile, "", log.LstdFlags)
	} else {
		l = logger
	}
	return
}

func CloseLog() {
	logFile.Close()
}

func getTime() (timeString string) {
	return time.Now().String()
}

func Println(msg string) {
	logger.SetPrefix("I")
	logger.Println(msg)
	fmt.Println(msg)
}

func Printf(format string, v ...interface{}) {
	logger.SetPrefix("I")
	logger.Printf(format, v)
	fmt.Printf(format, v)
}

func Errorf(format string, v ...interface{}) {
	logger.SetPrefix("E")
	logger.Printf(format, v)
	fmt.Printf(format, v)
}

func Error(msg string) {
	logger.SetPrefix("E")
	logger.Println(msg)
	fmt.Println(msg)
}
