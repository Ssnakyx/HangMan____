package HangMan_2023

import (
	"bufio"
	"math/rand"
	"os"
)

type HangManData struct {
	Word              string
	Found             []string
	AttemptsRemaining int
	GameStage         []string
}

var maxAttempts = 6

func IsWordGuessed(word string, found []string) bool {
	for _, letter := range word {
		if !Contains(found, string(letter)) {
			return false
		}
	}
	return true
}

func Contains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func ReadWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func GetRandomWordFromList(words []string) string {
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}

func GetUserGuess() string {
	var guess string
	return guess
}

func DisplayGameStatus(word string, found []string, attemptsRemaining int, stage []string) {
	DisplayWord(word, found)
}

func DisplayWord(word string, found []string) string {
	displayedWord := ""
	for _, letter := range word {
		if Contains(found, string(letter)) {
			displayedWord += string(letter) + " "
		} else {
			displayedWord += "_ "
		}
	}
	return displayedWord
}
