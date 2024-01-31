package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	id int
	p  *TreeNode
	vl *TreeNode
	vr *TreeNode
}

func createTree(n int, lim int, pn *TreeNode, linkN map[int]*TreeNode) (rootNode *TreeNode) {
	t := TreeNode{
		id: n,
		p:  pn,
	}
	linkN[n] = &t
	if n*2 <= lim {
		t.vl = createTree(n*2, lim, &t, linkN)
	} else {
		t.vl = nil
	}
	if n*2+1 <= lim {
		t.vr = createTree(n*2+1, lim, &t, linkN)
	} else {
		t.vr = nil
	}
	return &t
}

func outputTreeValues(root *TreeNode, res *[]int) {
	//*res = append(*res, root.id)
	if root.vl != nil {
		outputTreeValues(root.vl, res)
	}
	*res = append(*res, root.id)
	if root.vr != nil {
		outputTreeValues(root.vr, res)
	}

}

func searchByIdNew(i int, linkN map[int]*TreeNode) *TreeNode {
	return linkN[i]
}

func swap(v *TreeNode, head *TreeNode) *TreeNode {
	if v == head {
		return head
	}
	tempP := v.p
	if tempP != nil {
		tempPP := v.p.p
		if tempP.vl == v {
			tempP.vl = v.vl
			v.p = tempPP
			if tempPP != nil {
				if tempPP.vl == tempP {
					tempPP.vl = v
				} else {
					tempPP.vr = v
				}
			} else {
				head = v
			}
			v.vl = tempP
			tempP.p = v

			if tempP.vl != nil {
				tempP.vl.p = tempP
			}
		} else {
			tempP.vr = v.vr

			v.p = tempPP
			if tempPP != nil {
				if tempPP.vl == tempP {
					tempPP.vl = v
				} else {
					tempPP.vr = v
				}
			} else {
				head = v
			}

			v.vr = tempP
			tempP.p = v

			if tempP.vr != nil {
				tempP.vr.p = tempP
			}
		}
	}
	return head
}

func main() {
	newarray := []int{}
	newlink := &newarray

	linkNodes := make(map[int]*TreeNode)

	var sl []string
	var slPt []int
	temp := ""
	n, q := 0, 0
	tempId := 0
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
	sl = strings.SplitN(temp, " ", 2)
	sl[1] = strings.TrimSpace(sl[1])
	n, _ = strconv.Atoi(sl[0])
	q, _ = strconv.Atoi(sl[1])

	temp, err = reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sl = strings.SplitN(temp, " ", q)
	sl[q-1] = strings.TrimSpace(sl[q-1])
	for i := 0; i < q; i++ {
		tempId, _ = strconv.Atoi(sl[i])
		slPt = append(slPt, tempId)
	}
	head := createTree(1, n, nil, linkNodes)

	for _, v := range slPt {
		head = swap((searchByIdNew(v, linkNodes)), head)
	}
	outputTreeValues(head, newlink)
	for _, v := range newarray {
		temp = strconv.Itoa(v)
		writer.WriteString(temp)
		writer.WriteString(" ")
	}
	writer.WriteString("")
	writer.Flush()
}
