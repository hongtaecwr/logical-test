package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter pin code: ")
	scanner.Scan()

	ok := validatePinCode(scanner.Text())
	if !ok {
		log.Println(scanner.Text(), ok)
	} else {
		log.Println(scanner.Text(), ok)
	}
}

func validatePinCode(input string) bool {
	regexpCompile := regexp.MustCompile(`\d{6,}`)
	if !regexpCompile.MatchString(input) {
		return false
	} else {
		regexpDuplicate := regexp.MustCompile(`\d`)
		result := regexpDuplicate.FindAllString(input, -1)

		var intSlice []int
		for _, i := range result {
			j, _ := strconv.Atoi(i)
			intSlice = append(intSlice, j)
		}

		diffSequence := 0
		for i := 1; i < len(intSlice); i++ {
			if intSlice[i] > intSlice[i-1] {
				diff := intSlice[i] - intSlice[i-1]
				if diff == 1 {
					diffSequence++
					if diffSequence == 2 {
						return false
					}
				} else {
					diffSequence = 0
				}
			} else if intSlice[i] < intSlice[i-1] {
				diff := intSlice[i-1] - intSlice[i]
				if diff == 1 {
					diffSequence++
					log.Println(diffSequence)
					if diffSequence == 2 {
						return false
					}
				} else {
					diffSequence = 0
				}
			}
		}
		dupCount := 0
		for i := 0; i < len(intSlice); i++ {
			for j := i + 1; j < len(input); j++ {
				if input[i] == input[j] {
					dupCount++
					if dupCount > 2 {
						return false
					}
				}
			}
		}

	}
	return true
}
