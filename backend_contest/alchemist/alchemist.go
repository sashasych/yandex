package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type request struct {
	numOfA    int
	numOfB    int
	potionNum int
}

type potion struct {
	consist []int
	needA   int64
	needB   int64
	ok      bool
}

func createPotion(arr []int) potion {
	return potion{
		consist: arr,
		needA:   0,
		needB:   0,
		ok:      true,
	}
}

func createRequest(a int, b int, pNum int) request {
	return request{
		numOfA:    a,
		numOfB:    b,
		potionNum: pNum,
	}
}

func calcAandB(id int, potions *[]potion, check *map[int]bool) (int64, int64, bool) {
	potion := (*potions)[id]
	var a, b int64
	a, b = 0, 0

	//проверяем не является ли рецепт ложным
	//если рецепт уже посчитан, выводим результат
	if potion.ok && (potion.needA > 0 || potion.needB > 0) {
		return potion.needA, potion.needB, true
	} else if !potion.ok {
		return 0, 0, false
	}
	// проверяем нет ли замыкания
	// если есть возвращаем ложь
	if _, try := (*check)[id]; try {
		potion.ok = false       // !!!
		(*potions)[id] = potion // !!!
		return 0, 0, false
	} else {
		(*check)[id] = true
	}
	// рекурсивно считаем состав
	// если есть замыкание, все рецепты будут ложными
	for _, val := range potion.consist {

		if val == 1 {
			a += 1
		} else if val == 2 {
			b += 1
		} else if val > 2 { // && val <= 500000
			tempA, tempB, ok := calcAandB(val, potions, check)
			if ok {
				a += tempA
				b += tempB
			} else {
				potion.ok = false
				(*potions)[id] = potion
				return 0, 0, false
			}
		}
		/*
			if a > 100000000 || b > 100000000 {
				potion.ok = false
				potions[id] = potion
				return 0, 0, false
			}
		*/
	}

	// !!!
	/*
		if !potion.ok {
			(*potions)[id] = potion
			return 0, 0, false
		}
	*/
	potion.needA = a
	potion.needB = b
	(*potions)[id] = potion
	return a, b, true
}

func main() {
	var temp string
	var n, q, l int
	sum := 0
	var slStr []string

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
	writer := bufio.NewWriterSize(fileOutput, 1024*1024*32)

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	n, _ = strconv.Atoi(temp)
	datPotion := make([]potion, n+1)
	for i := 3; i <= n; i++ {
		temp, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		slStr = strings.Split(temp, " ")
		l = len(slStr)
		for i := 0; i < l; i++ {
			slStr[i] = strings.TrimSpace(slStr[i])
		}
		k, err := strconv.Atoi(slStr[0])
		if err != nil {
			panic(err)
		}
		slInt := make([]int, k)
		sum += k
		for i := 0; i < k; i++ {
			slInt[i], err = strconv.Atoi(slStr[i+1])
			if err != nil {
				panic(err)
			}
		}
		datPotion[i] = createPotion(slInt)
	}
	if sum > 1000000 {
		for _, val := range datPotion {
			val.ok = false
		}
	}
	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	q, err = strconv.Atoi(temp)
	if err != nil {
		panic(err)
	}
	slInt := make([]int, 3)
	datRequest := make([]request, q)
	for i := 0; i < q; i++ {
		temp, _ = reader.ReadString('\n')
		slStr = strings.SplitAfterN(temp, " ", 3)
		slStr[2] = strings.TrimSpace(slStr[2])
		slStr[1] = strings.TrimSpace(slStr[1])
		slStr[0] = strings.TrimSpace(slStr[0])
		slInt[0], err = strconv.Atoi(slStr[0])
		if err != nil {
			panic(err)
		}
		slInt[1], err = strconv.Atoi(slStr[1])
		if err != nil {
			panic(err)
		}
		slInt[2], err = strconv.Atoi(slStr[2])
		if err != nil {
			panic(err)
		}
		datRequest[i] = createRequest(slInt[0], slInt[1], slInt[2])
	}

	if sum <= 1000000 {
		mid := n / 2
		for id := mid; id >= 3; id-- {
			check := make(map[int]bool)
			_, _, _ = calcAandB(id, &datPotion, &check)
		}
		for id := mid + 1; id <= n; id++ {
			check := make(map[int]bool)
			_, _, _ = calcAandB(id, &datPotion, &check)
		}
	}

	l = len(datRequest)
	for id := 0; id < l; id++ {
		tempReq := datRequest[id]
		tempPotion := datPotion[tempReq.potionNum]
		if int64(tempReq.numOfA) >= tempPotion.needA && int64(tempReq.numOfB) >= tempPotion.needB && tempPotion.ok {
			writer.WriteRune('1')
		} else {
			writer.WriteRune('0')
		}

	}
	writer.Flush()
}
