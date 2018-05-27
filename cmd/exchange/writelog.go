package main

import (
	"fmt"
	"os"
)

const (
	WAL_DIRECTORY = "./wal"
)

type WriteLog struct {
	exchange string
	fillLog  *os.File
}

func NewWriteLog(exchange string) *WriteLog {
	filePath := WAL_DIRECTORY + exchange
	var log *os.File
	//Check if the directory is setup or need to be setup
	if _, err := os.Stat(WAL_DIRECTORY); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := os.Mkdir(WAL_DIRECTORY, 0777)
		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(WAL_DIRECTORY + "/" + exchange); os.IsNotExist(err) {
		// path/to/whatever does not exist
		f, err := os.Create(filePath)
		log = f
		if err != nil {
			panic(err)
		}
	}

	return &WriteLog{
		exchange: exchange,
		fillLog:  log,
	}
}

//Close closes the file that backs a writelog
func (l *WriteLog) Close() {
	l.fillLog.Close()
}

func (l *WriteLog) logFill(fill Fill) {
	l.fillLog.WriteString(fmt.Sprintf("%+v \n", fill))
}
