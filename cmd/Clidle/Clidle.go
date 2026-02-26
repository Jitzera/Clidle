package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

var wordBank = [...]string{"aback", "abase", "abate", "abbey", "abbot", "abhor"}

var correctWord = ""
var runedCorrect = []rune("")
var guessedWord = ""
var runedGuess = []rune("")

//var triesLeft = 5
//var correctLetters = 0

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
	debugWord := "water"
	runedCorrect = []rune(debugWord)
}

// TODO only allow words in wordbank
func insertScanner() {
	keyScan := ""
	fmt.Scanln(&keyScan)
	guessedWord = keyScan
	runedGuess = []rune(guessedWord)

	if len(guessedWord) == 5 {
		wordVerify()
	} else {
		println("Invalid input.")
	}
}

var inventoryLength = 5
var guessSliceInventory = []rune("")
var correctSliceInventory = []rune("")

func wordVerify() {
	guessSliceInventory = slices.Clone(runedGuess)
	correctSliceInventory = slices.Clone(runedCorrect)
	fmt.Print(runedCorrect)

	inventoryLength = len(guessSliceInventory)
	for i := range 5 {
		if runedGuess[i] == runedCorrect[i] {
			guessSliceInventory = slices.Delete(guessSliceInventory, 0, 1) //Err
			inventoryLength = len(guessSliceInventory)
			correctSliceInventory = slices.Delete(correctSliceInventory, 0, 1) //Err
			//fmt.Print(correctSliceInventory)
			print("=")
			//fmt.Print(guessSliceInventory)
		} else {
			//incorrectLetters++
			yellowChecker()
		}

	}

}

// Add correct deletion position
func yellowChecker() {
	inventoryLength = len(correctSliceInventory)
	for i := range inventoryLength {
		if correctSliceInventory[i] == guessSliceInventory[0] {
			//add sensor to see what letter is recognized
			print("real")
			fmt.Print(correctSliceInventory)
			guessSliceInventory = slices.Delete(guessSliceInventory, 0, 1)
			correctSliceInventory = slices.Delete(correctSliceInventory, 0, 1)
			fmt.Print(correctSliceInventory)
			break
		}
	}
}

func wrongPrint() {
	//print("-")
	//wordVerify()
	// drop inventory
}

/*
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
*/
func quit() {
}
