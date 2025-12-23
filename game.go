package main

import "fmt"

func main() { // Changed from game() to main()
	min := 1
	max := 100
	var s string

	fmt.Println("Think of a number between 1 and 100")

	for {
		guess := (min + max) / 2
		fmt.Printf("Is your number %d? (h/l/c): ", guess)
		fmt.Scanln(&s)

		if s == "h" {
			max = guess
		} else if s == "l" {
			min = guess
		} else if s == "c" {
			fmt.Println("Yay! I guessed your number!")
			break
		} else {
			fmt.Println("Please enter h, l, or c")
		}
	}
}
