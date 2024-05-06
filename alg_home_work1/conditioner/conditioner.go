package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tempRoom, tempCond int
	var tempData []string
	temp := ""
	// открываем файл на чтение
	fileInput, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	// отложенное закрытие
	defer fileInput.Close()
	reader := bufio.NewReaderSize(fileInput, 1024*1024*128)

	// открываем файл на запись
	fileOutput, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// отложенное закрытие
	defer fileOutput.Close()
	writer := bufio.NewWriterSize(fileOutput, 1024*1024*8)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//temp = strings.TrimSpace(temp)

	tempData = strings.SplitN(temp, " ", 2)
	// разбили строку и удаляем пустые символы из последней строки
	tempData[1] = strings.TrimSpace(tempData[1])
	tempRoom, _ = strconv.Atoi(tempData[0])
	tempCond, _ = strconv.Atoi(tempData[1])

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	switch temp {
	case "heat":
		if tempRoom <= tempCond {
			writer.WriteString(strconv.FormatInt(int64(tempCond), 10))
		} else {
			writer.WriteString(strconv.FormatInt(int64(tempRoom), 10))
		}
	case "freeze":
		if tempRoom >= tempCond {
			writer.WriteString(strconv.FormatInt(int64(tempCond), 10))
		} else {
			writer.WriteString(strconv.FormatInt(int64(tempRoom), 10))
		}
	case "auto":
		writer.WriteString(strconv.FormatInt(int64(tempCond), 10))
	case "fan":
		writer.WriteString(strconv.FormatInt(int64(tempRoom), 10))
	}
	writer.Flush()
}
