package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	amount := parseInt(os.Args[1])
	denominationString := os.Args[2]
	fmt.Println(MakeChange(amount, denominationString))
}

// MakeChange makes change based on amount and available denomination
func MakeChange(amount int, denominationString string) int {
	denominationArray := strings.Split(denominationString, ",")
	denomination := make(map[int]bool)
	for i := 0; i < len(denominationArray); i++ {
		d := parseInt(denominationArray[i])
		denomination[d] = true
	}

	changeArray := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		var min = amount + 1
		for d := range denomination {
			if i-d >= 0 {
				min = minInt(min, changeArray[i-d]+1)
			}
		}
		changeArray[i] = min
	}

	return changeArray[amount]
}

func parseInt(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return output
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
