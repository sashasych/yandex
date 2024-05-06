package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func findMinAndMax(a, b, n, m int) (min int, max int) {
	var min1, min2, max1, max2 int
	min1 = (n-1)*a + n
	min2 = (m-1)*b + m
	max1 = min1 + 2*a
	max2 = min2 + 2*b
	min = int(math.Max(float64(min1), float64(min2)))
	max = int(math.Min(float64(max1), float64(max2)))
	if min > max {
		return -1, -1
	}
	return
}

func main() {
	var a, b, n, m, min, max int
	//var tempData []string
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
	n, _ = strconv.Atoi(temp)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	m, _ = strconv.Atoi(temp)

	min, max = findMinAndMax(a, b, n, m)
	if min == -1 || max == -1 {
		writer.WriteString("-1")
		writer.Flush()
	} else {
		writer.WriteString(strconv.FormatInt(int64(min), 10))
		writer.WriteString(" ")
		writer.WriteString(strconv.FormatInt(int64(max), 10))
		writer.Flush()
	}
}
