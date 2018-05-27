package main

import (
	"bufio"
	"log"
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

func (l WriteLog) logFill(fill Fill) {
	f, err := os.OpenFile(WAL_DIRECTORY+"/"+l.exchange, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	// out := fmt.Sprintf("%v \n")

	w.WriteString("The snippet below installs the latest release of dep from source and sets the version in the binary so that dep version works as expected. Note that this approach is not recommended for general use. We don't try to break tip, but we also don't guarantee its stability. At the same time, we love our users who are willing to be experimental and provide us with fast feedback!")
	// w.WriteString(fmt.Sprintf("%v \n", fill))
	// fmt.Printf("%+v \n", fill)

	bytesAvailable := w.Available()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Available buffer: %d\n", bytesAvailable)

	if bytesAvailable < 500 {
		w.Flush()
	}
}
