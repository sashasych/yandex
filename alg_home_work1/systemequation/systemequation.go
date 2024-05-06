package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cramerSolve(a, b, c, d, e, f float64) string {
	str := ""
	result := ""
	var x, y, kk, bb float64
	// (5)
	if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 && f == 0 {
		return "5"
	}
	if (a == 0 && b == 0 && e != 0) || (c == 0 && d == 0 && f != 0) {
		return "0"
	}
	determinant := a*d - b*c
	deterX := e*d - f*b
	deterY := a*f - c*e
	//fmt.Println(deterX, deterY)
	// (2)
	if determinant != 0 {
		x = deterX / determinant
		y = deterY / determinant
		result = "2 "
		//str = strconv.FormatFloat()
		str = fmt.Sprintf("%.5f ", x)
		result += str
		str = fmt.Sprintf("%.5f", y)
		result += str
		return result
	} else {
		if deterX != 0 || deterY != 0 {
			return "0"
		} else {
			// y0 x - любое (4)
			if a == 0 && c == 0 {
				result = "4 "
				if b != 0 {
					y = e / b
				} else {
					y = f / d
				}

				str = fmt.Sprintf("%.5f", y)
				result += str
				return result
				//  (3) x0 y - любое
			} else if b == 0 && d == 0 {
				result = "3 "
				if a != 0 {
					x = e / a
				} else {
					x = f / c
				}
				str = fmt.Sprintf("%.5f", x)
				result += str
				return result
				// (1)
			} else {
				if b == 0 {
					//if a == 0 && b == 0 {
					kk = -c / d
					bb = f / d
				} else {
					kk = -a / b
					bb = e / b
				}
				result = "1 "
				str = fmt.Sprintf("%.5f ", kk)
				result += str
				str = fmt.Sprintf("%.5f", bb)
				result += str
				return result
			}
		}
	}

}

func main() {
	var a, b, c, d, e, f float64
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
	a, _ = strconv.ParseFloat(temp, 64)
	//a, _ = strconv.Atoi(temp)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	b, _ = strconv.ParseFloat(temp, 64)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	c, _ = strconv.ParseFloat(temp, 64)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	d, _ = strconv.ParseFloat(temp, 64)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	e, _ = strconv.ParseFloat(temp, 64)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	f, _ = strconv.ParseFloat(temp, 64)

	writer.WriteString(cramerSolve(a, b, c, d, e, f))
	writer.Flush()
}
