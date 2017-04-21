package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func makeSequence(numDigits int) string {
	array := make([]string, numDigits+1)
	for i := range array {
		array[i] = strconv.Itoa(i + 1)
	}

	return strings.Join(array, ",")
}

var makeChangeTests = []struct {
	amount             int
	denominationString string
	expected           int
}{
	// AmountEqualToDenomination
	{1, "1,3,5", 1},
	{3, "1,3,5", 1},
	{5, "1,3,5", 1},

	// CanadianDenominationAndValidAmount
	{4, "1,5,10,25,100,200", 4},
	{9, "1,5,10,25,100,200", 5},
	{42, "1,5,10,25,100,200", 5},
	{50, "1,5,10,25,100,200", 2},
	{68, "1,5,10,25,100,200", 7},
	{230, "1,5,10,25,100,200", 3},
	{250, "1,5,10,25,100,200", 3},
	{330, "1,5,10,25,100,200", 4},
	{500, "1,5,10,25,100,200", 3},

	// IrregularDenominationAndValidAmount
	{6, "1,3,4", 2},
	{7, "1,3,4,5", 2},
	{194, "1,3,5,10,12,50,100", 6},

	// LargeAmount
	{9999999, "1", 9999999},
}

func TestMakeChange(t *testing.T) {
	for _, tt := range makeChangeTests {
		fmt.Printf("%+v\n", tt)

		start := time.Now()
		actual := MakeChange(tt.amount, tt.denominationString)
		elapsed := time.Since(start)
		fmt.Println("MakeChange took", elapsed)

		if actual != tt.expected {
			t.Errorf("MakeChange(%d, %s): expected %d, actual %d", tt.amount, tt.denominationString, tt.expected, actual)
		}

		fmt.Println()
	}
}

func TestMakeChangeWithLargeNumberOfCoins(t *testing.T) {
	expected := 1

	fmt.Println("Making denomination")
	start := time.Now()
	denomination := makeSequence(999999)
	elapsed := time.Since(start)
	fmt.Println("Making denomination took", elapsed)

	fmt.Println("Running test")
	start = time.Now()
	actual := MakeChange(1, denomination)
	elapsed = time.Since(start)
	fmt.Println("MakeChangeWithLargeNumberOfCoins took", elapsed)

	if actual != expected {
		t.Errorf("MakeChangeWithLargeNumberOfCoins: expected %d, actual %d", expected, 1)
	}

	fmt.Println()
}
