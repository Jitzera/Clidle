package main

import "fmt"

var yesPrint = false
var correctWord = ""
var runedCorrect = []rune("")
var guessedWord = ""
var runedGuess = []rune("")
var triesLeft = 5
var correctLetters = 0

var secondGuess = rune(0)

func main() {
	println("6 Tries. (-) Means incorrect, (+) Out of position, (=) Correct. Type guess:")
	wordRNG()
	wordScanner()
}

// TODO add an array of words and an RNG to select one at random
// ability to read wordbank from a different file
func wordRNG() {
	//[...]wordBank{"water", "quake"}
	correctWord = "quake"
	runedCorrect = []rune(correctWord)
}

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
		println("")
		println("win")
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
		quit()
	}
}

func quit() {

}

//store incorrect letter in a bank?
// would require storing the print until the end to overwite

func wordVerifier() {
	letterCount := 0
	for letterCount < 5 {

		if runedGuess[letterCount] == runedCorrect[letterCount] {
			correctLetters++
			print("=")
		} else {
			secondGuess = runedGuess[letterCount]
			otherPositionChecker()
			if yesPrint == true {
				print("+")
			} else {
				print("-")
			}
		}
		letterCount++
	}

}

func otherPositionChecker() {
	secCount := 0
	yesPrint = false
	for secCount < 5 {
		if secondGuess == runedCorrect[secCount] {
			yesPrint = true
		}
		secCount++
	}

}
