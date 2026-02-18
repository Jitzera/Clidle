package main

import (
	"fmt"
	"math/rand/v2"
)

var correctSlice = make([]int, 5)
var guessSlice = make([]int, 5)

var wordBank = [...]string{"aback", "abase", "abate", "abbey", "abbot", "abhor"}

var letterInWord = false
var letterInWordCount = 0
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
	debugWord := "xxxxb"
	runedCorrect = []rune(debugWord)
}

// TODO only allow words in wordbank
func wordScanner() {
	keyScan := ""
	fmt.Scanln(&keyScan)
	guessedWord = keyScan
	runedGuess = []rune(guessedWord)

	if len(guessedWord) == 5 {
		//wordVeri//make var for each letter in word?fier()
		wordChecker()
	} else {
		println("Invalid input.")
		wordScanner()
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
		wordScanner()
	} else {
		fmt.Printf("\nLost. The word was %v.", correctWord)
		quit()
	}
}

func quit() {
}

// TODO add repeated letter detection
func wordVerifier() {
	letterIndex := 0
	for letterIndex < 5 {

		if runedGuess[letterIndex] == runedCorrect[letterIndex] {
			correctLetters++
			print("=")
			//make slice add guessed letters into it to prevent duplication
		} else {
			incorrectLetter = runedGuess[letterIndex]
			otherPositionChecker()
			if letterInWord == true {
				print("+")
				letterInWord = false
			} else {
				print("-")
			}
		}
		letterIndex++
	}

	gameState()
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

// make var for each letter in word?
// func to add a letter to a variable
// loops over all variables to see if it exists
// if yes add to it once
// if not add to an empty
// on check same in reverse but subtract each

func wordChecker() {

	for i := range 5 {
		if runedGuess[i] == runedCorrect[i] {
			//print("real")
			fmt.Print(runedGuess[i])
		}
	}

}
