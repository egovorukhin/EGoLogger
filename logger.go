package egologger

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
	system   string
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
		system:   runtime.GOOS,
	}
}

// SetFileSize Устанавливаем размер файлов
func (l Logger) SetFileSize(filesize int) Logger {
	l.fileSize = filesize
	return l
}

// SetLogPath Кастомный путь
func (l Logger) SetLogPath(path string, isAddPath bool) Logger {
	if isAddPath {
		l.filePath = filepath.Join(l.filePath, path)
	} else {
		l.filePath = path
	}
	return l
}

// Info Пишем префикс - "Info"
func (l Logger) Info(message interface{}) {

	if l.system == "windows" {
		logger := log.New(&l.buf, fmt.Sprintf("Info: %s - ", l.funcName), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix)
		err := logger.Output(2, fmt.Sprint(message))
		if err != nil {
			log.Fatal(err)
		}
		err = l.save()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// linux
	log.Println(message)
}

// Info Пишем префикс - "Info"
func Info(funcName interface{}, filename string, message interface{}) {
	New(funcName, filename).Info(message)
}

// Error Пишем префикс - "Error"
func (l Logger) Error(message interface{}) {

	if message == nil {
		return
	}

	if l.system == "windows" {
		logger := log.New(&l.buf, fmt.Sprintf("Error: %s - ", l.funcName), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix|log.Llongfile)
		err := logger.Output(2, fmt.Sprint(message))
		if err != nil {
			log.Fatal(err)
		}
		err = l.save()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// linux
	log.Println(message)
}

func Error(funcName interface{}, filename string, message interface{}) {
	if message == nil {
		return
	}
	New(funcName, filename).Error(message)
}
