package api

import (
	"fmt"
	"log"
	"marketplace/internal/api/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Router struct {
	handler handlers.Handlers
}

func NewRouters(handler handlers.Handlers) Router {
	return Router{
		handler: handler,
	}

}

func (r *Router) Start() {
	router := mux.NewRouter()

	// User endpoints
	router.HandleFunc("/api/users", r.handler.CreateUserHandler).Methods("POST")

	// Provider endpoints
	router.HandleFunc("/api/providers", r.handler.CreateProviderHandler).Methods("POST")
	router.HandleFunc("/api/getAllProviders", r.handler.GetAllProvidersHandler).Methods("GET")

	// Skill endpoints
	router.HandleFunc("/api/skills", r.handler.CreateSkillHandler).Methods("POST")
	router.HandleFunc("/api/skills", r.handler.UpdateSkillHandler).Methods("PUT")

	// Offer endpoints
	router.HandleFunc("/api/offer", r.handler.MakeOfferHandler).Methods("POST")
	router.HandleFunc("/api/offer", r.handler.UpdateOfferStatusHandler).Methods("PUT")

	// Task endpoints
	router.HandleFunc("/api/tasks", r.handler.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks", r.handler.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/getAllTasks", r.handler.GetAllTasksHandler).Methods("GET")
	router.HandleFunc("/api/task/progress", r.handler.UpdateTaskProgressHandler).Methods("PUT")
	router.HandleFunc("/api/task/status", r.handler.UpdateTaskCompletionStatusHandler).Methods("PUT")

	http.Handle("/", router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
