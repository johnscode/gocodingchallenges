package uniquecombos

import (
	"slices"
)

func FindUniqueCombinations(str string) []string {
	// Convert string to slice of runes
	runes := []rune(str)

	// Store unique combos
	var combos = []string{}

	// Generate combos recursively
	var searchForCombos func(start int, currentCombo []rune)
	searchForCombos = func(start int, currentCombo []rune) {
		comboAsStr := string(currentCombo)
		// Skip duplicates and empty string
		if !slices.Contains(combos, comboAsStr) && len(comboAsStr) > 0 {
			// Add current currentCombo to the result
			combos = append(combos, string(currentCombo))
		}

		// Generate combos for remaining characters
		for i := start; i < len(runes); i++ {
			currentCombo = append(currentCombo, runes[i])
			searchForCombos(i+1, currentCombo)
			currentCombo = currentCombo[:len(currentCombo)-1]
		}
	}

	searchForCombos(0, []rune{})
	return combos
}
