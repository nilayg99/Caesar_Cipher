package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func play_again() {
	fmt.Println("Want to play again ?")
	fmt.Printf("Type 'Y' for yes or 'N' for No : ")
	play_again_input := ""
	_, err := fmt.Scanln(&play_again_input)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		play_again()
	} else {
		user_input_lower := strings.ToLower(play_again_input)
		if user_input_lower == "y" {
			main()

		} else if user_input_lower == "n" {
			fmt.Printf("Thanks for encrypting with us .")
			os.Exit(0)

		} else {
			fmt.Println("Error reading input: ", err)
			play_again()

		}
	}
}

func main() {
	fmt.Printf("Enter 'E' to encrypt and 'D' Decrypt : ")
	user_input := ""
	_, err := fmt.Scanln(&user_input)
	if err != nil {
		fmt.Println("Error reading input : ", err)
		main()

	} else {

		user_input_lower := strings.ToLower(user_input)
		if user_input_lower == "e" {
			fmt.Println("E")
			play_again()

		} else if user_input_lower == "d" {
			fmt.Println("D")
			play_again()

		} else {
			fmt.Println("Error")
			main()

		}

	}

}
