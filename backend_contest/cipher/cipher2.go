package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type personData struct {
	surname string
	name    string
	midName string
	day     int
	month   int
	year    int
}

func createPersonData(sn string, n string, mn string, d int, m int, y int) personData {
	return personData{
		surname: sn,
		name:    n,
		midName: mn,
		day:     d,
		month:   m,
		year:    y,
	}
}

func getStringCipher(p personData) string {
	count := 0
	unicRune := make(map[rune]bool)
	d, m := p.day, p.month
	fio := p.name + p.surname + p.midName
	for _, letter := range fio {
		if _, ok := unicRune[letter]; !ok {
			unicRune[letter] = true
			count++
		}
	}
	sumDayMonthDigit := d%10 + d/10%10 + m%10 + m/10%10
	nAlphabetFirst := int(p.surname[0]) - 64
	final := count + sumDayMonthDigit*64 + nAlphabetFirst*256
	//fmt.Println(sumDayMonthDigit, y, nAlphabetFirst, final)
	//i, _ := strconv.ParseInt(final, 10, 64)
	answer := strconv.FormatInt(int64(final), 16)
	l := len(answer)
	if len(answer) < 3 {
		for i := 0; i < 3-l; i++ {
			answer = "0" + answer
		}
	} else {
		answer = answer[len(answer)-3:]
	}
	answer = strings.ToUpper(answer)
	return answer
}

func main() {

	var day, month, year, n int
	temp, answer := "", ""
	var tempData []string
	fileInput, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer fileInput.Close()
	reader := bufio.NewReaderSize(fileInput, 1024*1024*128)

	fileOutput, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileOutput.Close()
	writer := bufio.NewWriterSize(fileOutput, 1024*1024*8)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	n, _ = strconv.Atoi(temp)
	data := make([]personData, n)

	for i := 0; i < n; i++ {
		temp, _ = reader.ReadString('\n')
		tempData = strings.SplitN(temp, ",", 6)
		tempData[5] = strings.TrimSpace(tempData[5])
		day, _ = strconv.Atoi(tempData[3])
		month, _ = strconv.Atoi(tempData[4])
		year, _ = strconv.Atoi(tempData[5])
		data[i] = createPersonData(tempData[0], tempData[1], tempData[2], day, month, year)
	}

	for _, value := range data {
		answer = getStringCipher(value)
		writer.WriteString(answer)
		writer.WriteString(" ")
	}
	writer.Flush()
}
