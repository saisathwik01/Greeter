package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "greeter",
		Short: "Greeter is a basic CLI",
		Long:  "Greeter is a friendly command-line application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Greeter !!")
		},
	}

	greetCmd := &cobra.Command{
		Use:   "greet [name]",
		Short: "Greet the user",
		Long:  "Greet the user with a friendly message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			message, err := GetGreetingFromAPI(name)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println(message)
		},
	}

	rootCmd.AddCommand(greetCmd)

	go server() // Start the server in a separate goroutine

	// Print the URL
	fmt.Println("Server running at: http://localhost:8080/greet?name=YourName")

	// Block forever to keep the server running
	select {}
}
