package main

import (
	"bytes"
	"os"
)

const (
	WAL_DIRECTORY = "./wal"
)

type WriteLog struct {
	exchange string
}

func NewWriteLog(exchange string) *WriteLog {
	filePath := WAL_DIRECTORY + "/" + exchange
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
		_, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
	}
	return &WriteLog{
		exchange: exchange,
	}
}

func (l WriteLog) logFills(fills []Fill) {
	f, err := os.OpenFile(WAL_DIRECTORY+"/"+l.exchange, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// out := fmt.Sprintf("%v \n")
	var buffer bytes.Buffer

	for _, fill := range fills {
		buffer.WriteString(fill.Json() + "\n")
	}

	f.WriteString(buffer.String())
}

func (l WriteLog) logFill(fill Fill) {
	f, err := os.OpenFile(WAL_DIRECTORY+"/"+l.exchange, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// out := fmt.Sprintf("%v \n")

	f.WriteString(fill.Json() + "\n")
}
