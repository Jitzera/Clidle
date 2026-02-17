package main

import (
	"fmt"
	"math/rand/v2"
)

var wordBank = [...]string{"aback", "abase", "abate", "abbey", "abbot", "abhor"}

var letterInWord = false
var correctWord = ""
var runedCorrect = []rune("")
var guessedWord = ""
var runedGuess = []rune("")
var triesLeft = 5
var correctLetters = 0
var incorrectLetter = rune(0)

func main() {
	println("6 Tries. (-) Means incorrect, (+) Out of position, (=) Correct. Type guess:")
	wordRNG()
	wordScanner()
}

// ability to read wordbank from a different file
func wordRNG() {
	wordCount := len(wordBank)
	randIndex := rand.IntN(wordCount)
	correctWord = wordBank[randIndex]
	runedCorrect = []rune(correctWord)
}

// TODO only allow words in wordbank
func wordScanner() {
	keyScan := ""
	fmt.Scanln(&keyScan)
	guessedWord = keyScan
	runedGuess = []rune(guessedWord)

	if len(guessedWord) == 5 {
		wordVerifier()
	} else {
		println("Invalid input.")
	}
	gameState()
}

// TODO add timer and how many tries it took
func gameState() {
	if correctLetters == 5 {
		fmt.Println("\nwin")
	} else {
		correctLetters = 0
		livesRemaining()
	}
}
func livesRemaining() {
	if triesLeft != 0 {
		triesLeft--
		println("")
		wordScanner()
	} else {
		fmt.Printf("\nLost. The word was %v.", correctWord)
		quit()
	}
}

func quit() {
}

func wordVerifier() {
	letterIndex := 0
	for letterIndex < 5 {

		if runedGuess[letterIndex] == runedCorrect[letterIndex] {
			correctLetters++
			print("=")
		} else {
			incorrectLetter = runedGuess[letterIndex]
			otherPositionChecker()
			if letterInWord == true {
				print("+")
			} else {
				print("-")
			}
		}
		letterIndex++
	}

}

func otherPositionChecker() {
	Count := 0
	letterInWord = false
	for Count < 5 {
		if incorrectLetter == runedCorrect[Count] {
			letterInWord = true
		}
		Count++
	}

}
