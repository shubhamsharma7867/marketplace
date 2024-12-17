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

// getAllTaskCmd represents the getAllTask command
var getAllTaskCmd = &cobra.Command{
	Use:   "getAllTask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getAllTasks()
	},
}

func init() {
	rootCmd.AddCommand(getAllTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAllTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAllTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAllTasks() {
	cli := http.Client{}
	var tasks []models.Task
	url := "http://localhost:8080/api/getAllTasks"
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
	err = json.Unmarshal(bodyBytes, &tasks)
	if err != nil {
		log.Printf("failed to unmarshal response body : Error %v", err)
		return
	}

	fmt.Println(len(tasks))
}
