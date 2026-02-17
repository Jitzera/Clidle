package main

import "fmt"

var correctWord = ""
var runedCorrect = []rune("")
var guessedWord = ""
var runedGuess = []rune("")

func main() {
	wordRNG()
	wordScanner()
}

// TODO add an array of words and an RNG to select one at random
// ability to read wordbank from a different file
func wordRNG() {
	//[...]wordBank{"water", "quake"}
	correctWord = "water"
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
}

//store incorrect letter in a bank?
// would require storing the print until the end to overwite

func wordVerifier() {
	count := 0
	for count < 5 {

		if runedGuess[count] == runedCorrect[count] {
			print("=")
		} else {
			print("-")
		}
		count++
	}

}
