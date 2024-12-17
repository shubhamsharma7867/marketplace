package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"marketplace/internal/models"
	"net/http"
)

func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	if user.UserType == "BUSINESS" && !h.requestValidator.Validate(user.CompanyFeilds) {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", fmt.Errorf("failed to validate company feilds")), http.StatusBadRequest)
		return
	}
	err := h.db.Insert(h.userTableName, user)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.userTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.userTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	err := h.db.Insert(h.taskTableName, task)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.taskTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.taskTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handlers) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	cond := models.SqlCondition{
		Condition: fmt.Sprintf("%s = ?", "task_id"),
		Values:    []interface{}{task.TaskID},
	}
	rowsUpdated, err := h.db.Updates(h.taskTableName, task, cond)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.taskTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.taskTableName), http.StatusInternalServerError)
		return
	}
	if rowsUpdated == 0 {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.taskTableName, fmt.Errorf("no rows updated for condition %v ", cond))
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.taskTableName), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func (h *Handlers) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	err := h.db.Get(h.taskTableName, &tasks, models.SqlCondition{
		Condition: "",
		Values:    []interface{}{},
	}, 50, 0, "")
	if err != nil {
		log.Printf("Error while getting rows in %s table. Error : %v \n", h.taskTableName, err)
		http.Error(w, fmt.Sprintf("Error while getting rows in %s table ", h.taskTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h *Handlers) UpdateTaskCompletionStatusHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	cond := models.SqlCondition{
		Condition: fmt.Sprintf("%s = ?", "task_id"),
		Values:    []interface{}{task.TaskID},
	}
	rowsUpdated, err := h.db.Updates(h.taskTableName, task, cond)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.taskTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.taskTableName), http.StatusInternalServerError)
		return
	}
	if rowsUpdated == 0 {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.taskTableName, fmt.Errorf("no rows updated for condition %v ", cond))
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.taskTableName), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Task updated successfully"})
}

func (h *Handlers) UpdateOfferStatusHandler(w http.ResponseWriter, r *http.Request) {
	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	cond := models.SqlCondition{
		Condition: fmt.Sprintf("%s = ?", "offer_id"),
		Values:    []interface{}{offer.OfferID},
	}
	rowsUpdated, err := h.db.Updates(h.offerTableName, offer, cond)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.offerTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.offerTableName), http.StatusInternalServerError)
		return
	}
	if rowsUpdated == 0 {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.offerTableName, fmt.Errorf("no rows updated for condition %v ", cond))
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.offerTableName), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Offer status updated successfully"})
}
