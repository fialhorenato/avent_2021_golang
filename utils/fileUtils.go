package utils

import (
	"bufio"
	"container/list"
	"log"
	"os"
)

func GetContentAsString(fileName string) *list.List {
	list := list.New()

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		list.PushBack(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	return list
}
