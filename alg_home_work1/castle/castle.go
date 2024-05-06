package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkHole(a, b, c, d, e int) string {
	var smallerSide, biggerSide, smallSide, midleSide int
	if d < e {
		smallerSide = d
		biggerSide = e
	} else {
		smallerSide = e
		biggerSide = d
	}
	if a <= b && a <= c {
		smallSide = a
		if b <= c {
			midleSide = b
		} else {
			midleSide = c
		}
	}
	if b <= a && b <= c {
		smallSide = b
		if a <= c {
			midleSide = a
		} else {
			midleSide = c
		}
	}
	if c <= a && c <= b {
		smallSide = c
		if b <= a {
			midleSide = b
		} else {
			midleSide = a
		}
	}
	if smallSide <= smallerSide && midleSide <= biggerSide {
		return "YES"
	}
	return "NO"
}

func main() {
	var a, b, c, d, e int
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
	c, _ = strconv.Atoi(temp)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	d, _ = strconv.Atoi(temp)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	e, _ = strconv.Atoi(temp)

	writer.WriteString(checkHole(a, b, c, d, e))
	writer.Flush()
}
