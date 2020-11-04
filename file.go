package egologger

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

func (l Logger) save() error {

	//Возвращаем путь к директории с логами
	if _, err := os.Stat(l.filePath); os.IsNotExist(err) {
		err = os.MkdirAll(l.filePath, 0777)
		if err != nil {
			return err
		}
	}

	if l.fileName == "" {
		l.fileName = "unknown"
	}

	//Формируем полный путь к файлу логов
	l.fileName = filepath.Join(l.filePath, l.fileName+".log")

	//Проверяем путь на корректность
	info, err := os.Stat(l.fileName)
	if !os.IsNotExist(err) {

		if info == nil {
			return errors.New("Неверный формат пути!")
		}

		//Проверяем размер файла и удаляем если превышает установленный размер
		if info.Size() > int64(l.fileSize)*1024*1024 {
			err = os.Remove(l.fileName)
			if err != nil {
				return err
			}
		}
	}

	//Используем mutex за нормальную конкуренцию за память
	var mutex sync.Mutex
	ch := make(chan error)
	go l.write(ch, &mutex)

	return <-ch
}

func (l Logger) write(ch chan error, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	var err error
	//Открываем файл и раздаем права
	file, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		ch <- err
		return
	}

	//Пишем в файл данные
	_, err = file.Write(l.buf.Bytes())
	if err != nil {
		ch <- err
		return
	}

	//err отправляем в канал
	ch <- nil
}
