package main

import (
	"fmt"
	"sort"
	"strconv"
)

type romanNum struct {
	value int
	symbol string
}

func main() {
	s := intToRoman(3657)
	fmt.Println(s)
}

func intToRoman(num int) string {
	romanNums := []romanNum{
		{1, "I"},
		{5, "V"},
		{10, "X"},
		{50, "L"},
		{100, "C"},
		{500, "D"},
		{1000, "M"},
	}
	specialRomanNums := []romanNum{
		{4, "IV"},
		{9, "IX"},
		{40, "XL"},
		{90, "XC"},
		{400, "CD"},
		{900, "CM"},
	}
	romanStarterValues := []int{1, 10, 100, 1000}
	numStr := strconv.Itoa(num)
	numComposition := getNumComposition(numStr, romanStarterValues)
	
	var allLowerCompositions []int
	for _, composition := range(numComposition) {
		allLowerCompositions = append(allLowerCompositions, getNumLowerComposition(composition, romanNums, specialRomanNums)...)
	}

	romanNumStr := ""
	for _, lowerComposition := range(allLowerCompositions) {
		romanNumStr = romanNumStr + lowerCompositionToRoman(lowerComposition, append(romanNums, specialRomanNums...))
	}

	return romanNumStr
}

func getNumComposition(numStr string, romanStarterValues []int) []int {
	var numComposition []int

	for i := 0; i < len(numStr); i++ {
		for j := 0; j < len(romanStarterValues); j++ {
			currentNum, _ := strconv.Atoi(numStr[i:])
			ratio := currentNum / romanStarterValues[j]
			numFromIndex, _ := strconv.Atoi(string([]rune(numStr)[i]))

			if ratio == numFromIndex {
				composition := ratio * romanStarterValues[j]
				numComposition = append(numComposition, composition)
			} else {
				continue
			}
		}
	}

	return numComposition
}

func getNumLowerComposition(composition int, romanNums []romanNum, specialRomanNums []romanNum) []int {
	var lowerCompositions []int

	for _, specialRomanNum := range(specialRomanNums) {
		if composition == specialRomanNum.value {
			lowerCompositions = append(lowerCompositions, composition)
			return lowerCompositions
		}
	}

	for _, currentRomanNum := range(romanNums) {
		currentVal := currentRomanNum.value
		currentValLength := len(strconv.Itoa(currentRomanNum.value))
		compositionLength := len(strconv.Itoa(composition))

		if  compositionLength < currentValLength {
			break
		} else if compositionLength > currentValLength {
			continue
		}

		if composition == currentVal {
			lowerCompositions = append(lowerCompositions, currentVal)
			break
		}

		if compositionLength == currentValLength && composition > currentVal {
			factor := composition / currentVal
			
			if factor == 5 {
				continue
			}
	
			if factor > 5 {
				factor = factor - 5
			}
	
			for j := 0; j < factor; j++ {
				lowerCompositions = append(lowerCompositions, currentVal)
			}
		}
	}

	sort.Slice(lowerCompositions, func(i, j int) bool {
		return lowerCompositions[i] > lowerCompositions[j]
	})

	return lowerCompositions
}

func lowerCompositionToRoman(lowerComposition int, romanNums []romanNum) string {
	for _, currentRomanNum := range(romanNums) {
		if lowerComposition == currentRomanNum.value {
			return currentRomanNum.symbol
		}
	}

	return ""
}