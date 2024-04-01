/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setuserCmd represents the setuser command
var setuserCmd = &cobra.Command{
	Use:   "setuser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setUsername(args[0])
	},
}

func setUsername(username string) {
	userID := viper.Get("user_id")
	if userID == "" {
		log.Fatalln("User ID not found. Please run 'start' command first.")
	}

	fmt.Println("Setting username for user ID:", userID)

	requestData := map[string]string{
		"name": username,
	}

	payload, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalln("Failed to marshal request data:", err)
	}

	// TODO change hardcoded URL to env or yml
	url := fmt.Sprintf("http://localhost:8080/session/set-name/%s", userID)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to set username:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Failed to set username. Server returned status:", resp.Status)
	}

	fmt.Println("Username set successfully.")
}

func init() {
	rootCmd.AddCommand(setuserCmd)
}
