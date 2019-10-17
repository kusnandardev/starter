package logging

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// Logger :
type Logger struct {
	UUID string `json:"uuid,omitempty"`
}

// Debug :
func (l *Logger) Debug(v ...interface{}) {
	l.setUserLogPrefix(DEBUG)
	log.Println(v...)
	logger.Println(v...)
}

// Info :
func (l *Logger) Info(v ...interface{}) {
	l.setUserLogPrefix(INFO)
	log.Println(v...)
	logger.Println(v...)
}

// Warn :
func (l *Logger) Warn(v ...interface{}) {
	l.setUserLogPrefix(WARNING)
	log.Println(v...)
	logger.Println(v...)
}

// Error :
func (l *Logger) Error(v ...interface{}) {
	l.setUserLogPrefix(ERROR)
	log.Println(v...)
	logger.Println(v...)
}

// Fatal :
func (l *Logger) Fatal(v ...interface{}) {
	l.setUserLogPrefix(FATAL)
	log.Println(v...)
	logger.Fatalln(v...)
}

func (l *Logger) setUserLogPrefix(level Level) {
	function, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		s := strings.Split(runtime.FuncForPC(function).Name(), ".")
		_, fn := s[0], s[1]
		logPrefix = fmt.Sprintf("[%s][%s][%s][%s:%d]", levelFlags[level], l.UUID, fn, filepath.Base(file), line)
		eFlag = levelFlags[level]
		eFunc = fn
		eFile = filepath.Base(file)
		eLine = line
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
