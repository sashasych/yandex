package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func getSides() {}

func findMax(n1 int, n2 int) int {
	if n1 >= n2 {
		return n1
	} else {
		return n2
	}
}

func main() {
	var x1, x2, y1, y2, s, t1, t2, r1, r2, tempS int
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
	tempData = strings.SplitN(temp, " ", 4)
	// разбили строку и удаляем пустые символы из последней строки
	tempData[3] = strings.TrimSpace(tempData[3])
	x1, _ = strconv.Atoi(tempData[0])
	x2, _ = strconv.Atoi(tempData[1])
	y1, _ = strconv.Atoi(tempData[2])
	y2, _ = strconv.Atoi(tempData[3])

	t1 = x1 + y1
	t2 = findMax(x2, y2)
	s = t1 * t2
	tempS = s
	r1, r2 = t1, t2

	t1 = x1 + y2
	t2 = findMax(x2, y1)
	s = t1 * t2
	if s < tempS {
		tempS = s
		r1, r2 = t1, t2
	}

	t1 = x2 + y1
	t2 = findMax(x1, y2)
	s = t1 * t2
	if s < tempS {
		tempS = s
		r1, r2 = t1, t2
	}

	t1 = x2 + y2
	t2 = findMax(x1, y1)
	s = t1 * t2
	if s < tempS {
		tempS = s
		r1, r2 = t1, t2
	}

	writer.WriteString(strconv.FormatInt(int64(r1), 10))
	writer.WriteString(" ")
	writer.WriteString(strconv.FormatInt(int64(r2), 10))
	writer.Flush()
}
