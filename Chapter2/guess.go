package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	success := false
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	for guesses := 10; guesses > 0; guesses-- {
		fmt.Println("\nYou have", guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("YOU GUESSED IT!!")
			success = true
			break
		} else if guess > target {
			fmt.Println("You guessed too high")
		} else {
			fmt.Println("You guessed too low")
		}
	}
	if !success {
		fmt.Println("You LOSE! The correct number was", target)
	}
}
