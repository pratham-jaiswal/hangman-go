package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var wordList = []string{
	"rollsroyce",
	"bentley",
	"lamborghini",
	"ferrari",
	"porsche",
	"koenigsegg",
	"mercedes",
	"bmw",
	"audi",
	"jaguar",
	"lexus",
	"maserati",
	"bugatti",
	"mclaren",
	"pagani",
	"tesla",
	"cadillac",
	"lincoln",
	"infiniti",
	"acura",
}
var chosenWord string
var guessWord []rune
var attempts int

func main() {
	rand.Seed(time.Now().UnixNano())
	chosenWord = wordList[rand.Intn(len(wordList))]

	guessWord = make([]rune, len(chosenWord))
	for i := range guessWord {
		guessWord[i] = '_'
	}

	attempts = int(float64(len(chosenWord)) * 1.5)

	reader := bufio.NewReader(os.Stdin)
	for attempts > 0 {
		fmt.Printf("Guess the word: %s\n", string(guessWord))
		fmt.Printf("Attempts remaining: %d\n", attempts)
		fmt.Print("Enter a letter: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		found := false
		for i, c := range chosenWord {
			if rune(input[0]) == c {
				guessWord[i] = c
				found = true
			}
		}

		if !found {
			attempts--
		}

		if string(guessWord) == chosenWord {
			fmt.Printf("Congratulations! You've guessed the word: %s\n", chosenWord)
			return
		}
	}

	fmt.Printf("Sorry, you've run out of attempts. The word was: %s\n", chosenWord)
}
