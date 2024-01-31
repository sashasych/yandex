package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	day    int
	hour   int
	minute int
	status string
}

func createEvent(days int, hours int, minutes int, statusIn string) Event {
	return Event{
		day:    days,
		hour:   hours,
		minute: minutes,
		status: statusIn,
	}
}

func calcFlightTime(slice []Event) (time string) {
	temp := 0
	//test := 0
	for i := 0; i < len(slice)-1; i++ {
		//test = 1440*(5+3) + 60*15 + 25
		if slice[i].status == "A" || slice[i].status == "B" {
			temp += (slice[i+1].day-slice[i].day)*1440 + (slice[i+1].hour-slice[i].hour)*60 + (slice[i+1].minute - slice[i].minute)
		}
	}
	//fmt.Println(test)
	return strconv.Itoa(temp)
}

func main() {

	//t := time.Now()
	var n, day, hour, minute, id int
	var statusIn string
	var splitStrings []string
	answer := ""
	temp := ""
	rockets := make(map[int][]Event)
	idsOfRockets := make([]int, 0)
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

	temp, _ = reader.ReadString('\n')
	//temp = strings.TrimRight(temp, "")
	temp = strings.TrimSpace(temp)
	n, _ = strconv.Atoi(temp)
	//fmt.Println(n)
	//n, _ = strconv.ParseInt(temp, 10, 0)
	//splitStrings = strings.Split(temp, " ")
	//fmt.Fscanln(reader, &n)

	//tempStrings := make([]string, 0, n)
	/*
		for i := 0; i < n; i++ {
			temp, _ = reader.ReadString('\n')
			tempStrings = append(tempStrings, temp)
		}
	*/
	for i := 0; i < n; i++ {
		temp, _ = reader.ReadString('\n')
		//temp = strings.TrimRight(temp, "\n")
		splitStrings = strings.SplitN(temp, " ", 5)
		day, _ = strconv.Atoi(splitStrings[0])
		hour, _ = strconv.Atoi(splitStrings[1])
		minute, _ = strconv.Atoi(splitStrings[2])
		id, _ = strconv.Atoi(splitStrings[3])
		statusIn = strings.TrimSpace(splitStrings[4])
		//fmt.Fscanln(reader, &day, &hour, &minute, &id, &statusIn)
		rockets[id] = append(rockets[id], createEvent(day, hour, minute, statusIn))
		//idsOfRockets = append(idsOfRockets, id)
		/*
			if _, isOk := rockets[id]; isOk {
				rockets[id] = append(rockets[id], createEvent(day, hour, minute, statusIn))
			} else {
				rockets[id] = make([]Event, 0)
				rockets[id] = append(rockets[id], createEvent(day, hour, minute, statusIn))
				idsOfRockets = append(idsOfRockets, id)
			}
		*/
	}
	for id = range rockets {
		idsOfRockets = append(idsOfRockets, id)
	}
	//idsOfRockets = append(idsOfRockets, id)

	sort.Slice(idsOfRockets, func(i, j int) bool {
		return idsOfRockets[i] < idsOfRockets[j]
	})
	lastID := len(idsOfRockets) - 1
	for id, value := range idsOfRockets {
		temp := rockets[value]
		sort.Slice(temp, func(i, j int) bool {
			if temp[i].day < temp[j].day {
				return true
			} else if temp[i].day > temp[j].day {
				return false
			} else {
				if temp[i].hour < temp[j].hour {
					return true
				} else if temp[i].hour > temp[j].hour {
					return false
				} else {
					if temp[i].minute < temp[j].minute {
						return true
					} else if temp[i].minute > temp[j].minute {
						return false
					} else {
						return false
					}
				}
			}
		})
		//answer = calcFlightTime(temp)
		if id == lastID {
			answer = calcFlightTime(temp)
		} else {
			answer = calcFlightTime(temp) + " "
		}
		writer.WriteString(answer)
	}
	/*
				rockets[value].type SortBy []Type

				func (a SortBy) Len() int           { return len(a) }
				func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
				func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

		}
		for i := 0; i < 10000; i++ {
			writer.WriteString("134353\n")
			//answer += "134353\n"
		}
	*/
	//answer = strings.TrimRight(answer, " ")
	//writer.WriteString(answer)
	//writer.WriteString("\n")
	writer.Flush()
	//end := time.Since(t)
	//fmt.Println(end)
}
