package logger

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

type Logger struct {
	funcName string
	fileName string
	buf      bytes.Buffer
	fileSize int
	filePath string
}

func New(funcName interface{}, filename string) Logger {

	//Имя функции
	fName := ""
	if funcName != nil {
		fName = runtime.FuncForPC(reflect.ValueOf(funcName).Pointer()).Name()
	}

	//Путь до директории с логами
	app, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fPath := filepath.Join(filepath.Dir(app), "logs")

	return Logger{
		funcName: fName,
		fileName: filename,
		filePath: fPath,
		fileSize: 10,
	}
}

//Устанавливаем размер файлов
func (l Logger) SetFileSize(filesize int) Logger {
	l.fileSize = filesize
	return l
}

//Кастомный путь
func (l Logger) SetLogPath(path string, isAddPath bool) Logger {
	if isAddPath {
		l.filePath = filepath.Join(l.filePath, path)
	} else {
		l.filePath = path
	}
	return l
}

func (logger Logger) Info(message interface{}) {

	l := log.New(&logger.buf, fmt.Sprintf("Info: %s - ", logger.funcName), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
	err := l.Output(2, fmt.Sprint(message))
	if err != nil {
		log.Fatal(err)
	}
	err = logger.save()
	if err != nil {
		log.Fatal(err)
	}
}

func Info(funcName interface{}, filename string, message interface{}) {
	New(funcName, filename).Info(message)
}

func (logger Logger) Error(message interface{}) {

	if message == nil {
		return
	}

	l := log.New(&logger.buf, fmt.Sprintf("Error: %s - ", logger.funcName), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix|log.Llongfile)
	err := l.Output(2, fmt.Sprint(message))
	if err != nil {
		log.Fatal(err)
	}
	err = logger.save()
	if err != nil {
		log.Fatal(err)
	}
}

func Error(funcName interface{}, filename string, message interface{}) {
	if message == nil {
		return
	}
	New(funcName, filename).Error(message)
}
