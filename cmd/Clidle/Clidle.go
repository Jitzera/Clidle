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

func wordRNG() {
	correctWord = "water"
	runedCorrect = []rune(correctWord)
}

func wordScanner() {
	keyScan := ""
	fmt.Scanln(&keyScan)
	guessedWord = keyScan
	runedGuess = []rune(guessedWord)

	if len(guessedWord) == 5 {
		println(guessedWord)
		wordVerifier()
	} else {
		println("Invalid input.")
	}
}

func wordVerifier() {
	count := 0
	for count < 5 {
		//print(runedGuess[count])
		//print(runedCorrect[count])
		if runedGuess[count] == runedCorrect[count] {
			println("c")
		}
		count++
	}
	println("lamo")

}
