package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
)

// GeneratePassword generates a random password with the specified length
func GeneratePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]|:;<>,.?/~"
	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Print("-----------------------------------------------------------------\n")
	fmt.Printf("Welcome, This program is created by Jbearp\n")
	fmt.Println("Use the password generator to ensure your security")
	fmt.Print("-----------------------------------------------------------------\n")

	for {
		passwordLength := 12 // Default length
		password, err := GeneratePassword(passwordLength)
		if err != nil {
			fmt.Println("Error generating password:", err)
			os.Exit(1)
		}

		fmt.Printf("This is your password: %s\n", password)
		fmt.Print("-----------------------------------------------------------------\n")
		fmt.Print("Would you like to generate another password? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "y" {
			fmt.Printf("The Matrix is everywhere, It’s all around us. Logging off...\n")
			cmd := exec.Command("cmd", "/c", "exit") // This attempts to close the Command Prompt on Windows
			cmd.Stdout = os.Stdout
			cmd.Run()
			break
		}

		// Clear the screen before generating the next password
		clearScreen()
	}

	fmt.Println("End of transmission.")
}
