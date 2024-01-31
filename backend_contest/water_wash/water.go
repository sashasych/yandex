package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type order struct {
	start int
	end   int
	cost  int
	time  int
}

type request struct {
	start   int
	end     int
	reqType int
}

func createOrder(s int, e int, c int) order {
	return order{
		start: s,
		end:   e,
		cost:  c,
		time:  e - s,
	}
}

func createRequest(s int, e int, r int) request {
	return request{
		start:   s,
		end:     e,
		reqType: r,
	}
}

func summCostOrders(ord []order, start int, end int) (result int64) {
	//fmt.Println(start, end, "cost")
	//fmt.Println(st)
	result = 0
	/*
		if st != -1 {
			for i, e := st, req.end; i < len(ord) && ord[i].start <= e; i++ {
				result += int64(ord[i].cost)
			}
		}
	*/
	if start >= 0 && end >= 0 {
		if start > 0 {
			result = int64(ord[end].cost - ord[start-1].cost)
		} else {
			result = int64(ord[end].cost)
		}
	}

	/*
		if start >= 0 {
			for i := start; i <= end; i++ {
				result += int64(ord[i].cost)
			}
		}

		/*
			for _, val := range ord {
				if req.start <= val.start && req.end >= val.start {
					result += int64(val.cost)
				}
			}
	*/
	return result
}

func summTimeOrders(ord []order, start int, end int) (result int64) {
	//fmt.Println(st)
	//fmt.Println(start, end, "time")
	result = 0
	/*
		if st != -1 {
			for i, s := st, req.start; i >= 0 && s <= ord[i].end; i-- {
				result += int64(ord[i].time)
			}
		}
	*/

	if start >= 0 && end >= 0 {
		if start > 0 {
			result = int64(ord[end].time - ord[start-1].time)
		} else {
			result = int64(ord[end].time)
		}
	}
	/*
		if start >= 0 {
			for i := start; i <= end; i++ {
				result += int64(ord[i].time)
			}

		}

		/*
			for _, val := range ord {
				if req.start <= val.end && req.end >= val.end {
					result += int64(val.time)
				}
			}
	*/
	return result
}

func findIndex(ord []order, req request, t int) (start int, end int) {

	var mid int
	start, end = -1, -1
	l := len(ord) - 1
	if t == 1 {
		min := 0
		high := l
		reqSt := req.start
		for min <= high {
			mid = (min + high) / 2
			//fmt.Println(mid)
			if mid != 0 {
				if ord[mid].start >= reqSt && ord[mid-1].start < reqSt {
					start = mid
					break
				}
			} else {
				if ord[mid].start >= reqSt {
					start = mid
					break
				}
			}
			if ord[mid].start < reqSt {
				//fmt.Println("bbb")
				min = mid + 1
			} else {
				//fmt.Println("aaa")
				high = mid - 1
			}
		}
		if start >= 0 {
			min = start
			high = len(ord) - 1
			reqEnd := req.end
			//fmt.Println("hello")
			for min <= high {
				mid = (min + high) / 2
				//fmt.Println(mid)
				if mid != l {
					if ord[mid].start <= reqEnd && ord[mid+1].start > reqEnd {
						end = mid
						break
					}
				} else {
					if ord[mid].start <= reqEnd {
						end = mid
						break
					}
				}
				if ord[mid].start > reqEnd {
					high = mid - 1
				} else {
					min = mid + 1
				}
			}
		}
	} else { // для заказов которые закончились в промежутке
		min := 0
		high := l
		reqSt := req.start
		for min <= high {
			mid = (min + high) / 2
			//fmt.Println(mid)
			if mid != 0 {
				if ord[mid].end >= reqSt && ord[mid-1].end < reqSt {
					start = mid
					break
				}
			} else {
				if ord[mid].end >= reqSt {
					start = mid
					break
				}
			}
			if ord[mid].end < reqSt {
				min = mid + 1
			} else {
				high = mid - 1
			}
		}
		if start >= 0 {
			min = start
			high = len(ord) - 1
			reqEnd := req.end
			for min <= high {
				mid = (min + high) / 2
				if mid != l {
					if ord[mid].end <= reqEnd && ord[mid+1].end > reqEnd {
						end = mid
						break
					}
				} else {
					if ord[mid].end <= reqEnd {
						end = mid
						break
					}
				}
				if ord[mid].end > reqEnd {
					high = mid - 1
				} else {
					min = mid + 1
				}
			}
		}
	}

	return start, end
}

// сделать двоичный поиск

/*
	func findStartIndex(ord []order, target request) int {
		ind := -1
		if target.end >= ord[0].start {
			for i := 0; i < len(ord); i++ {
				if target.start <= ord[i].start && target.end >= ord[i].start {
					ind = i
					break
				}
			}
		}
		return ind
	}

	func findEndIndex(ord []order, target request) int {
		ind := -1
		l := len(ord) - 1
		if target.start <= ord[l].end {
			for i := l; i >= 0; i-- {
				if target.end >= ord[i].end && target.start <= ord[i].end {
					ind = i
					break
				}
			}
		}

		return ind
	}
*/
func main() {
	//datOrdersEnd := []order{}
	datOrdersStart := []order{}
	datRequest := []request{}
	temp := ""
	var n, q, start, end int
	slStr := make([]string, 3)
	slTime := make([]int, 3)
	//answer := []int{}
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
	for i := 0; i < n; i++ {
		temp, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		slStr = strings.SplitN(temp, " ", 3)
		slStr[2] = strings.TrimSpace(slStr[2])
		for i := range slTime {
			slTime[i], _ = strconv.Atoi(slStr[i])
		}

		datOrdersStart = append(datOrdersStart, createOrder(slTime[0], slTime[1], slTime[2]))
	}
	datOrdersEnd := make([]order, len(datOrdersStart))
	copy(datOrdersEnd, datOrdersStart)
	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	temp = strings.TrimSpace(temp)
	q, _ = strconv.Atoi(temp)
	for i := 0; i < q; i++ {
		temp, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		slStr = strings.SplitN(temp, " ", 3)
		slStr[2] = strings.TrimSpace(slStr[2])
		for i := range slTime {
			slTime[i], _ = strconv.Atoi(slStr[i])
		}
		datRequest = append(datRequest, createRequest(slTime[0], slTime[1], slTime[2]))
	}

	sort.Slice(datOrdersStart, func(i, j int) bool {
		return datOrdersStart[i].start < datOrdersStart[j].start
	})
	sort.Slice(datOrdersEnd, func(i, j int) bool {
		return datOrdersEnd[i].end < datOrdersEnd[j].end
	})
	l := len(datOrdersStart)
	for i := 1; i < l; i++ {
		datOrdersStart[i].cost = datOrdersStart[i].cost + datOrdersStart[i-1].cost
	}
	l = len(datOrdersEnd)
	for i := 1; i < l; i++ {
		datOrdersEnd[i].time = datOrdersEnd[i].time + datOrdersEnd[i-1].time
	}

	//fmt.Println(datOrdersStart)
	//fmt.Println(datOrdersEnd)
	//fmt.Println(datRequest)
	for _, val := range datRequest {
		//fmt.Println("hello")
		if val.reqType == 1 {
			start, end = findIndex(datOrdersStart, val, 1)
			writer.WriteString(strconv.FormatInt((summCostOrders(datOrdersStart, start, end)), 10))
			//writer.WriteString(strconv.Itoa(summCostOrders(datOrdersStart, v, ind)))
			writer.WriteString(" ")
		} else {
			start, end = findIndex(datOrdersEnd, val, 2)
			writer.WriteString(strconv.FormatInt((summTimeOrders(datOrdersEnd, start, end)), 10))
			//writer.WriteString(strconv.Itoa(summTimeOrders(datOrdersEnd, v, ind)))
			writer.WriteString(" ")
		} // доделать вывод и получение данных
	}
	//fmt.Println(datOrdersStart, datRequest, n, q)
	//writer.WriteString("hello")
	writer.Flush()
}
