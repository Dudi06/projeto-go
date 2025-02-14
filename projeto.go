package main

import "fmt"

func main() {
	fmt.Println("welcom to my quiz game!")
	fmt.Printf("Enter you name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	numQuestions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		score++
	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Println("How many cores does the Ryzen 9 3900x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Printf("You scored %v out of %v", score, numQuestions)
	percent := (float64(score) / float64(numQuestions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
