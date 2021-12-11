package main

import (
	b "adventofcode/utils"
	"fmt"
	"strconv"
)

func main() {
	list := b.GetContentAsString("./resources/day1.txt")

	var current int
	var increases int

	for temp := list.Front(); temp != nil; temp = temp.Next() {
		if current == 0 {
			str, _ := temp.Value.(string)
			current, _ = strconv.Atoi(str)
		} else {
			str, _ := temp.Value.(string)
			myInteger, _ := strconv.Atoi(str)
			if myInteger > current {
				increases++
			}
			current = myInteger
		}
	}

	fmt.Println(increases)
}
