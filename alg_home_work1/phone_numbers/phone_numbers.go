package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkExistNumber(num1 string, num2 string) string {
	ind1, ind2 := 0, 0
	var nums1, nums2 [11]rune
	var count int

	for _, val := range num1 {
		if val < '0' || val > '9' {
			continue
		} else {
			nums1[ind1] = val
			ind1++
			count++
		}
	}
	if count == 7 {
		for i, j := 6, 10; i >= 0; i-- {
			nums1[j] = nums1[i]
			j--
		}
		nums1[0] = '8'
		nums1[1] = '4'
		nums1[2] = '9'
		nums1[3] = '5'
	}
	count = 0
	for _, val := range num2 {
		if val < '0' || val > '9' {
			continue
		} else {
			nums2[ind2] = val
			ind2++
			count++
		}
	}
	if count == 7 {
		for i, j := 6, 10; i >= 0; i-- {
			nums2[j] = nums2[i]
			j--
		}
		nums2[0] = '8'
		nums2[1] = '4'
		nums2[2] = '9'
		nums2[3] = '5'
	}
	for i := 1; i <= 10; i++ {
		if nums1[i] != nums2[i] {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var numberToCheck string
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
	numberToCheck = temp

	for i := 0; i <= 2; i++ {
		temp, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		temp = strings.TrimSpace(temp)
		writer.WriteString(checkExistNumber(numberToCheck, temp))
		if i != 2 {
			writer.WriteString("\n")
		}
	}

	writer.Flush()
}
