package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	char_values = "'~`!@#$%^&*()+_-={}|[]:;,./<>?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	shift       = 3
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
func encrypt_text(text string, shift int) string {
	var result strings.Builder
	for _, char := range text {
		index := strings.IndexRune(char_values, char)
		if index != -1 {
			new_index := (index + shift) % len(char_values)
			result.WriteByte(char_values[new_index])
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()

}

func decrypt_text(text string, shift int) string {
	var result strings.Builder
	for _, char := range text {
		index := strings.IndexRune(char_values, char)
		if index != -1 {
			new_index := (index - shift + len(char_values)) % len(char_values)
			result.WriteByte(char_values[new_index])
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
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
			fmt.Println("Type the word you want to encrypt : ")
			decrypt_text_input := ""
			_, err := fmt.Scanln(&decrypt_text_input)
			if err != nil {
				fmt.Println("Error : ", err)
			}
			encrypted := encrypt_text(decrypt_text_input, shift)
			fmt.Println("Encrypted text:", encrypted)
			play_again()

		} else if user_input_lower == "d" {
			fmt.Println("Type the word you want to decrypt : ")
			decrypt_text_input := ""
			_, err := fmt.Scanln(&decrypt_text_input)
			if err != nil {
				fmt.Println("Error : ", err)
			}
			decrypted := decrypt_text(decrypt_text_input, shift)
			fmt.Println("Decrypted text:", decrypted)
			play_again()

		} else {
			fmt.Println("Invalid input.")
			main()

		}

	}

}
