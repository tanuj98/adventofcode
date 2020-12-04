package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//trees("trees.txt")
	//numberOfPasswords(0, 0)
	//fmt.Print(isValid(123444))
	validatePassport()
	// passwordstwo("passwords.txt")
}

func validatePassport() {
	file, err := os.Open("./passport.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	linesss := make([][]string, 0)
	for {
		lines, err := myReadLine(file, reader)
		if err != nil || len(lines) == 0 {
			break
		}

		fmt.Println(lines)
		linesss = append(linesss, lines)
	}

	toCheck := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	count := 0
	for _, line := range linesss {
		m := make(map[string]string)
		combined := ""
		for _, pair := range line {
			for _, x := range strings.Split(pair, " ") {
				combined = combined + " " + x
			}
		}


		for _, pair := range strings.Split(combined, " ") {
			//fmt.Println(pair)
			if len(strings.Split(pair, ":")) != 2 {
				continue
			}
			m[strings.Split(pair, ":")[0]] = strings.Split(pair, ":")[1]
		}

		fmt.Println(m)
		isValid := true
		for _, c := range toCheck {
			val, ok := m[c]
			if !ok {
				isValid = false
				break
			}

			switch c {
			case "byr":
				if len(val) == 4 {
					num, err := strconv.Atoi(val)
					if err != nil {
						isValid = false
					}

					if num < 1920 || num > 2002 {
						isValid = false
					}
				} else {
					isValid = false
				}
			case "iyr":
				if len(val) == 4 {
					num, err := strconv.Atoi(val)
					if err != nil {
						isValid = false
					}

					if num < 2010 || num > 2020 {
						isValid = false
					}
				} else {
					isValid = false
				}
			case "eyr":
				if len(val) == 4 {
					num, err := strconv.Atoi(val)
					if err != nil {
						isValid = false
					}

					if num < 2020 || num > 2030 {
						isValid = false
					}
				} else {
					isValid = false
				}
			case "hgt":

				if len(val) == 4 {
					num, err := strconv.Atoi(val[:2])
					if err != nil {
						isValid = false
					}

					if val[2:] != "in" {
						isValid = false
					}

					if num < 59 || num > 76 {
						isValid = false
					}
				} else if len(val) == 5 {
					num, err := strconv.Atoi(val[:3])
					if err != nil {
						isValid = false
					}

					if val[3:] != "cm" {
						isValid = false
					}

					if num < 150 || num > 193 {
						isValid = false
					}
				} else {
					isValid = false
				}
			case "ecl":
				if !(val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth") {
					isValid = false
				}
			case "pid":
				if _, err := strconv.Atoi(val); err != nil || len(val) != 9 {
					isValid = false
				}
			case "hcl":
				if !strings.HasPrefix(val, "#") {
					isValid = false
				}

				if len(val) != 7 {
					isValid = false
				}

				for _, r := range val {
					isValidChar := false
					if r >= 48 || r <= 57 {
						isValidChar = true
					}

					if r >= 97 && r <= 102 {
						isValidChar = true
					}

					isValid = isValid && isValidChar
				}
			}
		}

		if isValid {
			count++
		}
	}

	fmt.Println(count)
}

func myReadLine(file *os.File, reader *bufio.Reader) (lines []string, err error){
	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, err
}
