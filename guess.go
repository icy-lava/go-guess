package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	// "os"
)

func getNextGuess() int {
	for true {
		fmt.Print("Guess the number: ")
		var line string
		fmt.Scanln(&line)
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("That's not a number!")
			continue
		}
		return num
	}
	panic("unexpected loop break")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	minNum := 1
	maxNum := 100
	
	game:
	for true {
		actual := rand.Intn(maxNum - minNum + 1) + minNum

		fmt.Printf("I came up with a number between %d and %d\n\n", minNum, maxNum)
		guessCount := 0
		for true {
			guess := getNextGuess()
			guessCount++

			if guess < minNum || guess > maxNum {
				fmt.Printf("%d is not even within the range!", guess)
				continue
			}

			var response string
			if guess < actual {
				response = "too low"
			} else if guess > actual {
				response = "too high"
			} else {
				fmt.Printf("\n%d is my number, that's right!\n", guess)
				fmt.Printf("You guessed it in %d tries!\n\n", guessCount)
				
				for true {
					fmt.Print("Do you want to play again? (Y/n): ")
					
					var line string
					fmt.Scanln(&line)
					line = strings.Trim(strings.ToLower(line), " \t")
					
					yes := strings.Contains(line, "y")
					no := strings.Contains(line, "n")
				
					if (yes && !no) || (len(line) == 0) {
						fmt.Println("\nOK let's try again!")
						continue game
					} else if no && !yes {
						fmt.Println("\nOh OK. Maybe next time!")
						break game
					} else {
						fmt.Println("\nI don't know what you mean by that")
					}
				}
			}

			fmt.Printf("Your guess of %d is %s\n", guess, response)
		}
	}
}
