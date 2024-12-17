package main

import (
	"fmt"
	"marketplace/internal/api"
	"marketplace/internal/api/handlers"
	"marketplace/internal/db"
	"marketplace/internal/utils"
)

func main() {
	db, err := db.NewDb("root", "12345", "localhost", "test", 3306, "provider", "user", "skill", "task", "offer")
	if err != nil {
		fmt.Println("failed to create new db connections", err)
		return
	}
	validator := utils.Validator{}
	handler := handlers.NewHandlers(db, "provider", "skill", "task", "offer", "user", validator)
	router := api.NewRouters(handler)
	router.Start()

}
