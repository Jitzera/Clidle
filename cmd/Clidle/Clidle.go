package main

import (
	"fmt"
	"math/rand/v2"
)

var wordBank = [...]string{"aback", "abase", "abate", "abbey", "abbot", "abhor"}

var correctWord = ""
var runedCorrect = []rune("")
var guessedWord = ""
var runedGuess = []rune("")

var triesLeft = 5
var correctLetters = 0

func main() {
	println("6 Tries. (-) Means incorrect, (+) Out of position, (=) Correct. Type guess:")
	wordRNG()
	insertScanner()
}

// ability to read wordbank from a different file
func wordRNG() {
	wordCount := len(wordBank)
	randIndex := rand.IntN(wordCount)
	correctWord = wordBank[randIndex]
	debugWord := "xxxxb"
	runedCorrect = []rune(debugWord)
}

// TODO only allow words in wordbank
func insertScanner() {
	keyScan := ""
	fmt.Scanln(&keyScan)
	guessedWord = keyScan
	runedGuess = []rune(guessedWord)

	if len(guessedWord) == 5 {
		//wordVeri//make var for each letter in word?fier()
	} else {
		println("Invalid input.")
	}
}

func wordVerify() {
	sliceGuessInventory := runedGuess

	for i := range 5 {

	}

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
		fmt.Printf(" Tries left:%v", triesLeft)
		triesLeft--
		println("")
		insertScanner()
	} else {
		fmt.Printf("\nLost. The word was %v.", correctWord)
		quit()
	}
}

func quit() {
}
