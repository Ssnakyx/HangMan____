package HangMan_2023

import (
	"math/rand"
	"strings"
	"time"
)

func First() {
	var choice int

	var filename string
	switch choice {
	case 1:
		filename = "HangMan_2023/words.txt"
	case 2:
		filename = "HangMan_2023/words2.txt"
	case 3:
		filename = "HangMan_2023/words3.txt"
	}

	words, err := ReadWordsFromFile(filename)
	if err != nil {
	}

	rand.Seed(time.Now().UnixNano())
	word := GetRandomWordFromList(words)
	letter := make([]string, 0)
	maxAttempts := 6
	attempts := maxAttempts
	stage := []string{}
	DisplayGameStatuss(word, letter, attempts, stage)
}
func DisplayGameStatuss(word string, letter []string, attempts int, stage []string) {

	for attempts > 0 {
		guess := GetUserGuess()
		letter = append(letter, guess)

		guessedCorrectly := IsWordGuessed(word, letter)

		if strings.Contains(word, guess) {
		} else {
			if !guessedCorrectly {
				attempts--
			}
		}

		if guessedCorrectly {
			break
		}
	}

	if attempts == 0 {
	}
}
