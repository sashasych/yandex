package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcPandN(k1, m, k2, p2, n2 int) (p1 int, n1 int) {

	// этаж больше этажей в доме
	if m < n2 {
		return -1, -1
	}

	// 1- ый этаж первого подъезда два варианта
	if n2 == 1 && p2 == 1 {
		if k1 <= k2 {
			return 1, 1
		} else {
			if m == 1 {

				return 0, 1
			} else {
				numberOfFlatsIn := m * k2
				if k1 <= numberOfFlatsIn {
					return 1, 0
				} else {
					return 0, 0
				}
			}
		}
	}
	// все остальное (проверяем не ложь если кол квартир мин на этаж ниже не сходится ложь)

	// этаж относительно подъездов
	floorAll := n2 + m*(p2-1)
	// находим сред кол квартир на этаже (мин) не округленное
	temp := float64(k2) / float64((floorAll))
	// нашли мин кол квартир на этаже
	flatOnFloorMin := int(math.Ceil(temp))

	// (проверяем не ложь если кол квартир мин на этаж ниже не сходится ложь)
	if (floorAll-1)*flatOnFloorMin >= k2 {
		return -1, -1
	}

	// проверка, что квартира не та же
	if k2 == k1 {
		p1 = p2
		n1 = n2
		return
	}

	// дальше просчеты мин макс

	// создаем макс кол квартир на этаже
	flatOnFloorMax := 0

	// найдем макс кол квартир на этаже
	for i := 1; true; i++ {

		if (flatOnFloorMin+i)*(floorAll-1) >= k2 {
			break
		}
		flatOnFloorMax = flatOnFloorMin + i
	}
	//fmt.Println(flatOnFloorMax)
	floorFirstFlat := 0 // предположительный этаж первой квартиры
	if flatOnFloorMax == 0 {
		floorFirstFlat = int(math.Ceil(float64(k1) / float64(flatOnFloorMin)))
		// ищем этаж
		n1 = floorFirstFlat % m
		if n1 == 0 {
			n1 = m
		}
		p1 = int(math.Ceil(float64(floorFirstFlat) / float64(m)))
		return
	} else { // смотрим диапазон мин и макс
		// найдем подъезд

		floorFirstFlat = int(math.Ceil(float64(k1) / float64(flatOnFloorMin)))
		p1 = int(math.Ceil(float64(floorFirstFlat) / float64(m)))
		floorFirstFlat = int(math.Ceil(float64(k1) / float64(flatOnFloorMax)))
		temp := int(math.Ceil(float64(floorFirstFlat) / float64(m)))
		if temp != p1 {
			p1 = 0
		}

		// ищем этаж

		floorFirstFlat = int(math.Ceil(float64(k1) / float64(flatOnFloorMin)))
		//fmt.Println(floorFirstFlat)
		n1 = floorFirstFlat % m
		if n1 == 0 {
			n1 = m
		}
		//fmt.Println(n1)
		for i := flatOnFloorMax - flatOnFloorMin; i > 0; i-- {
			floorFirstFlat = int(math.Ceil(float64(k1) / float64(flatOnFloorMin+i)))
			temp = floorFirstFlat % m
			if temp == 0 {
				temp = m
			}
			//fmt.Println(temp)
			if temp != n1 {
				n1 = 0
				break
			}
		}
		return
	}
}

func main() {
	var k1, m, k2, p2, n2, p1, n1 int
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

	tempData = strings.SplitN(temp, " ", 5)
	// разбили строку и удаляем пустые символы из последней строки
	tempData[4] = strings.TrimSpace(tempData[4])
	k1, _ = strconv.Atoi(tempData[0])
	m, _ = strconv.Atoi(tempData[1])
	k2, _ = strconv.Atoi(tempData[2])
	p2, _ = strconv.Atoi(tempData[3])
	n2, _ = strconv.Atoi(tempData[4])

	p1, n1 = calcPandN(k1, m, k2, p2, n2)
	writer.WriteString(strconv.FormatInt(int64(p1), 10))
	writer.WriteString(" ")
	writer.WriteString(strconv.FormatInt(int64(n1), 10))
	writer.Flush()
}
