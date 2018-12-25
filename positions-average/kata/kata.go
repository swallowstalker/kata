package kata

import (
	"strings"
)

func PosAverage(s string) float64 {

	numbers := strings.Split(s, ", ")
	occurences := findOccurences(numbers)
	sameCombination := findCombination(occurences)
	totalCombination := int64(len(numbers[0])) * calculateCombination(int64(len(numbers)))

	return float64(sameCombination) / float64(totalCombination) * 100
}

func findOccurences(numbers []string) []int64 {
	var occurences []int64

	for i := 0; i < len(numbers[0]); i++ {
		lib := make(map[uint8]int64, len(numbers))

		for _, number := range numbers {
			lib[number[i]]++
		}

		for _, occurence := range lib {
			occurences = append(occurences, occurence)
		}

	}

	return occurences
}

func findCombination(lib []int64) int64 {
	var totalCombination int64

	for _, occurence := range lib {
		totalCombination += calculateCombination(occurence)
	}

	return totalCombination
}

func calculateCombination(occurence int64) int64 {
	if occurence < 2 {
		return 0
	}
	return occurence * (occurence - 1) / 2
}
