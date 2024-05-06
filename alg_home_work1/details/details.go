package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcGears(n, k, m int) int {
	mSum := 0
	if k < m {
		return mSum
	}

	if n/k >= 2 {
		count := 0
		plusDet := k / m
		back := k % m
		minusSplav := k - back
		count = (n - k) / minusSplav
		n = n - count*minusSplav
		mSum += count * plusDet
	}

	for n >= k {
		// вычитаем из сплава массу для заготовок
		n -= k
		// считаем детали из заготовок и добавляем
		mSum += k / m
		// возвращаем остатки заготовок в сплав
		n += k % m
	}
	return mSum
}

func main() {
	var n, k, m int
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

	tempData = strings.SplitN(temp, " ", 3)
	// разбили строку и удаляем пустые символы из последней строки
	tempData[2] = strings.TrimSpace(tempData[2])
	n, _ = strconv.Atoi(tempData[0])
	k, _ = strconv.Atoi(tempData[1])
	m, _ = strconv.Atoi(tempData[2])

	//fmt.Println(calcGears(n, k, m))
	writer.WriteString(strconv.FormatInt(int64((calcGears(n, k, m))), 10))
	writer.Flush()
}
