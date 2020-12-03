package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func passwords(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count, totalLines int
	for scanner.Scan() {
		totalLines++
		i := scanner.Text()
		split := strings.Split(i, ":")
		toCheck := split[1]
		condition := strings.Split(split[0], " ")
		toCheckChar := rune(condition[1][0])
		min, _ := strconv.Atoi(strings.Split(condition[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(condition[0], "-")[1])
		countMap := make(map[rune]int)
		for _, r := range toCheck {
			if _, ok := countMap[r]; !ok {
				countMap[r] = 0
			}

			countMap[r] = countMap[r] + 1
		}

		if countMap[toCheckChar] <= max && countMap[toCheckChar] >= min {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
	fmt.Println(totalLines)
}

func passwordstwo(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count, totalLines int
	for scanner.Scan() {
		totalLines++
		i := scanner.Text()
		split := strings.Split(i, ":")
		toCheck := split[1]
		condition := strings.Split(split[0], " ")
		toCheckChar := rune(condition[1][0])
		min, _ := strconv.Atoi(strings.Split(condition[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(condition[0], "-")[1])
		countMap := make(map[int]rune)
		for index, r := range toCheck {
			countMap[index] = r
		}

		fmt.Printf("%d %d %s %c \n", min, max, toCheck, toCheckChar)
		fmt.Printf("%c %c \n", countMap[min], countMap[max])
		fmt.Println("-----------------")
		pass := 0
		if countMap[min] == toCheckChar {
			pass++
		}

		if countMap[max] == toCheckChar {
			pass++
		}

		if pass == 1 {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
	fmt.Println(totalLines)
}
