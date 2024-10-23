package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// basic input
	// var userInput string
	// fmt.Print("Enter the text: ")
	// fmt.Scanln(&userInput)

	// input using bufio. basic input doesnt work well when there is spaces
	// create new bufio.NewReader
	input := bufio.NewReader(os.Stdin)

	// ReadLine can be used to clear any existing newline character from the input buffer before taking the actual user input
	// input.ReadLine()

	// reading the input
	fmt.Print("Enter value for input 2: ")
	userInput2, _ := input.ReadString('\n')

	// if userinput2 passed as it is it will pass \r - 13 in ascii and \n - 10 in ascii as well
	// if you input chama youll get -> [99 104 97 109 97 13 10]
	// to avoid passing \r and \n, following can be done
	// if you input chama youll get -> [99 104 97 109 97]
	userInput2 = strings.TrimSpace(userInput2)

	// counting letters of the string excluding spaces
	countLetters(userInput2)

	// couting words by iterating the string
	iterateStringOne(userInput2)

	iterateStringTwo(userInput2)

	countWordsRunes(userInput2)

	countWordsBySplit(userInput2)
}

// to count the letters of string just using len
func countLetters(userInput string) {
	fmt.Println("Counting Letters")
	fmt.Println(len(userInput) - countSpaces(userInput))
	fmt.Println("------------------")
}

func countSpaces(userInput string) int {
	count := 0
	for i := 0; i < len(userInput); i++ {
		if userInput[i] == ' ' {
			count++
		}
	}

	return count
}

func iterateStringOne(userInput string) {
	fmt.Println("by iterating the string and checking the character")
	count := 0
	for i := 0; i < len(userInput); i++ {
		// to print a character
		// fmt.Printf("%c", userInput[i])

		if userInput[i] == ' ' {
			count++
		}
	}
	fmt.Println(count + 1)
	fmt.Println("------------------")
}

func iterateStringTwo(userInput string) {
	fmt.Println("By itering the string using for-range loop")
	counter := 0
	for _, letter := range userInput {

		// byte value
		// fmt.Println(letter)

		// covent byte value to string
		// fmt.Println(string(letter))

		if letter == ' ' {
			counter++
		}
	}

	fmt.Println("By String iteration")
	fmt.Println(counter + 1)
	fmt.Println("------------------")
}

func countWordsRunes(userInput string) {
	fmt.Println("By converting to rune slice")
	// creating array of runes. runes are int32 values. so the array is array of int32
	runes := []rune(userInput)
	// fmt.Println("Runes: ", runes)

	// to get type of data reflect.TypeOf can be used
	// fmt.Println(reflect.TypeOf(runes))

	// iterating in the array of runes
	counter := 0
	for _, value := range runes {
		if value == 32 {
			counter++
		}
	}

	fmt.Print("Number of words: ", (counter + 1))
	fmt.Println("------------------")

}

func countWordsBySplit(userInput string) {
	fmt.Println("By spliting the string")
	words := strings.Fields(userInput)
	fmt.Println(words)
	fmt.Println(len(words))
	fmt.Println("------------------")
}
