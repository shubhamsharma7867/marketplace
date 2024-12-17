/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"marketplace/internal/models"
	"net/http"

	"github.com/spf13/cobra"
)

// getAllProvidersCmd represents the getAllProviders command
var getAllProvidersCmd = &cobra.Command{
	Use:   "getAllProviders",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getAllProviders()
	},
}

func init() {
	rootCmd.AddCommand(getAllProvidersCmd)
}

func getAllProviders() {
	cli := http.Client{}
	var providers []models.Provider
	url := "http://localhost:8080/api/getAllProviders"
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		log.Printf("failed to create http request : Error %v\n", err)
		return
	}

	res, err := cli.Do(req)
	if err != nil {
		log.Printf("failed to do get request : Error %v", err)
		return
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to do read response body : Error %v", err)
		return
	}
	err = json.Unmarshal(bodyBytes, &providers)
	if err != nil {
		log.Printf("failed to unmarshal response body : Error %v", err)
		return
	}
	fmt.Println(len(providers))
}
