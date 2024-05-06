package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solveEquation(a, b, c float64) string {
	var answer float64
	if c < 0 || (a == 0 && b < 0) {
		return "NO SOLUTION"
	} else if a == 0 {
		if math.Sqrt(b) == c {
			return "MANY SOLUTIONS"
		} else {
			return "NO SOLUTION"
		}
	} else {
		answer = (c*c - b) / a
	}
	_, frac := math.Modf(answer)
	if frac == 0 {
		return strconv.FormatInt(int64(int(answer)), 10)
	} else {
		return "NO SOLUTION"
	}
}

func main() {
	var a, b, c int
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

	temp = solveEquation(float64(a), float64(b), float64(c))
	writer.WriteString(temp)
	writer.Flush()
}
