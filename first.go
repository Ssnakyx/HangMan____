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

	words, err := readWordsFromFile(filename)
	if err != nil {
	}

	rand.Seed(time.Now().UnixNano())
	word := getRandomWordFromList(words)
	letter := make([]string, 0)
	maxAttempts := 6
	attempts := maxAttempts
	stage := []string{}

	displayGameStatus(word, letter, attempts, stage)

	for attempts > 0 {
		guess := getUserGuess()
		letter = append(letter, guess)

		guessedCorrectly := isWordGuessed(word, letter)

		if strings.Contains(word, guess) {
		} else {
			if !guessedCorrectly {
				attempts--
			}
		}

		displayGameStatus(word, letter, attempts, stage)

		if guessedCorrectly {
			break
		}
	}

	if attempts == 0 {
	}
}
