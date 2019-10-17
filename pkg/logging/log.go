package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"kusnandartoni/starter/pkg/file"
)

// Level :
type Level int

// F :
var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	eFlag      = ""
	eFunc      = ""
	eFile      = ""
	eLine      = -1
)

// DEBUG :
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

// Setup :
func Setup() {
	now := time.Now()
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)

	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	timeSpent := time.Since(now)
	log.Printf("Config logging is ready in %v", timeSpent)
}

// Debug :
func Debug(user string, v ...interface{}) {
	setPrefix(DEBUG)
	log.Println(v...)
	logger.Println(v...)
}

// Info :
func Info(user string, v ...interface{}) {
	setPrefix(INFO)
	log.Println(v...)
	logger.Println(v...)
}

// Warn :
func Warn(user string, v ...interface{}) {
	setPrefix(WARNING)
	log.Println(v...)
	logger.Println(v...)
}

// Error :
func Error(user string, v ...interface{}) {
	setPrefix(ERROR)
	log.Println(v...)
	logger.Println(v...)
}

// Fatal :
func Fatal(user string, v ...interface{}) {
	setPrefix(FATAL)
	log.Println(v...)
	logger.Fatalln(v...)
}

func setPrefix(level Level) {
	function, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		s := strings.Split(runtime.FuncForPC(function).Name(), ".")
		_, fn := s[0], s[1]
		logPrefix = fmt.Sprintf("[%s][SYS][%s][%s:%d]", levelFlags[level], fn, filepath.Base(file), line)
		eFlag = levelFlags[level]
		eFunc = fn
		eFile = filepath.Base(file)
		eLine = line
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
