# EGoLogger
Logger for golang projects
Логер для проектов написанных на Golang

#### Описание (Description)
Логер умеет писать в разные директории и ограничивать файлы для записи

#### Установка (Installation)
>go get github.com/egovorukhin/egologger

#### Использование (Usage)
>logPhones := New(nil, "phones").SetFileSize(1).SetLogPath("dir", true)\
>logPhones.Info("phone")\
>logPhones.Error("phone")\