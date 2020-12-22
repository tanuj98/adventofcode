package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay22(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ings, ings2 := readInput(t, "day22_example.txt")
		fmt.Println(ings, ings2)
		day22(ings, ings2)
		assert.Equal(t, false, true)
	})
	// t.Run("question", func(t *testing.T) {
	// 	ings, ings2 := readInput(t, "day22_actual.txt")
	// 	day22(ings, ings2)
	// })
}

func readInput(t *testing.T, filePath string) ([]int, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ings := make([]int, 0)
	toAppend := make([]int, 0)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(t, "Player 1") {
			continue
		}

		if strings.Contains(t, "Player 2") {
			for _, x := range toAppend {
				ings = append(ings, x)
			}

			toAppend = make([]int, 0)
			continue
		}

		num, _ := strconv.Atoi(t)
		toAppend = append(toAppend, num)

	}

	fmt.Println(ings)
	return ings[:len(ings)-1], toAppend
}


