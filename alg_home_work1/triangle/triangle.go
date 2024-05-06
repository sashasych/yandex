package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func triangleExist(a, b, c int) string {
	if (a+b > c) && (a+c > b) && (c+b > a) {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var a, b, c int
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

	//считываем стороны
	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	a, _ = strconv.Atoi(temp)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	b, _ = strconv.Atoi(temp)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	c, _ = strconv.Atoi(temp)

	// проверяем треугольник
	temp = triangleExist(a, b, c)
	// записываем результат
	writer.WriteString(temp)
	writer.Flush()
}
