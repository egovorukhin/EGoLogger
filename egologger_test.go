package egologger

import (
	"errors"
	"fmt"
	"testing"
)

type Main struct{}

func (m Main) GetError(text string) error {
	return errors.New(text)
}

func Test(t *testing.T) {

	fmt.Println("Старт")

	logPhones := New(nil, "phones").SetFileSize(1).SetLogPath("dir", true)
	logPhones.Info("телефон")
	logPhones.Error("телефон")

	log := New(Test, "phones")
	log.Error("Ошибка информации")
	log.Info("Старт")
	m := Main{}
	err := m.GetError("Ошибка")
	if err != nil {
		log.Error(err)
	}

	log.Info("Стоп")

	fmt.Println("Стоп")
	/*
		//В начале пишем имя приложения, это нужно для создания здиректории в Linux.
		//Для Windows данная процедура не нужна
		logger.SetAppName("EGoLogger")
		//Для изменения пути нужно изменить путь до директории с логами.
		//Директория logs будет создаваться в любом случае
		logger.SetLogPath("")
		//Редактируем размер файлов
		logger.SetFileSize(15)
		//Устанавливаем имя основного файла
		//logger.SetFileName("app")
		m := Main{}
		logger.Info("main", "Просто информация")
		logger.InfoFileName("main", "Просто информация", "main")
		err := m.GetError("Какая то ошибка!")
		if err != nil {
			logger.Error(m, err)
			logger.ErrorPrefix("main", err)
			logger.Trace(m, main, err)
			logger.Debug(m, main, err)
		}*/
}
