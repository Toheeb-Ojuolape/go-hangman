package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var pl = fmt.Println

/*
 +---+
 0   |
/|\  |
/ \  |
    ===
Secret word : M_____
Incorrect Guesses: A
Guess a Letter : Y

Sorry Your Dead! The word is MONKEY
Yes the Secret Word is MONKEY

Please Enter Only One Letter
Please Enter a Letter
Please Enter a Letter you haven't guessed
*/

var hangmanArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
	"ZODIAC", "ZOMBIE", "FLUFF",
}

var randWord string
var guessedLetters string
var correctLetter []string
var wrongGuesses []string

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordArr[rand.Intn(7)]
	correctLetter = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	pl(hangmanArr[len(wrongGuesses)])
	fmt.Print("Secret Word : ")
	for _, v := range correctLetter {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Print("\nIncorrect Guesses :")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	fmt.Println()
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nGuess a Letter : ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if len(guess) != 1 {
			fmt.Println("Please Enter Only One Letter")
		} else if !isLetter(guess) {
			fmt.Println("Please Enter a letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please enter a letter you haven't guessed")
		} else {
			return guess
		}
	}
}

func getAllCorrectLetterIndexes(theStr, subStr string) (indices []int) {
	if len(subStr) == 0 || (len(theStr) == 0) {
		return indices
	}
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllCorrectLetterIndexes(randWord, letter)
	for _, v := range indexMatches {
		correctLetter[v] = letter
	}
}

func sliceHasEmptys(theSlice []string) bool {
	for _, v := range theSlice {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func main() {
	getRandWord()
	for {
		showBoard()
		guess := getUserInput()
		//if correct letter is guessed
		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
			if sliceHasEmptys(correctLetter) {
				fmt.Println("More Letters to Guess")
			} else {
				fmt.Println("Yes the Secret Word is", randWord, "🎊")
				break
			}
		} else {
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)
			if len(wrongGuesses) >= 6 {
				fmt.Println("GAME OVER! The word is", randWord, "😥")
				break
			}
		}
	}
	// Show Gameboard

	// Get a letter from the user

	//A. If they guessed letter in word
	// Add to correctLetter
	//1. Are there more letters to guess?
	//2. If no more letters to guess, YOU WIN
	//B. If they guessed letter not in word
	//1. Add new letter to guessedLetters and wrongGuesses
	// 2. Check if they have run out of guesses
}
