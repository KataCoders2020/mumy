package main




func IsVowel (letter string) bool {
	vowels := [10] string {
		"a",
		"e",
		"i",
		"o",
		"u",
		"A",
		"E",
		"I",
		"O",
		"U",
	}

	for _, vowel := range vowels {
		if vowel == letter {
			return true
		}
	}
	return false
}

func VowelConverter (word string) string {
	return "Fred"
}
