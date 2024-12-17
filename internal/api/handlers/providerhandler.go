package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"marketplace/internal/models"
	"net/http"
)

func (h *Handlers) CreateProviderHandler(w http.ResponseWriter, r *http.Request) {
	var provider models.Provider
	if err := json.NewDecoder(r.Body).Decode(&provider); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	if provider.ProviderType == "BUSSINESS" && !h.requestValidator.Validate(provider.CompanyFeilds) {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", fmt.Errorf("failed to validate company feilds")), http.StatusBadRequest)
		return
	}
	err := h.db.Insert(h.providerTableName, provider)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.providerTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.providerTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(provider)
}

func (h *Handlers) GetAllProvidersHandler(w http.ResponseWriter, r *http.Request) {
	var providers []models.Provider

	err := h.db.Get(h.providerTableName, &providers, models.SqlCondition{
		Condition: "",
		Values:    []interface{}{},
	}, 50, 0, "")
	if err != nil {
		log.Printf("Error while getting rows in %s table. Error : %v \n", h.providerTableName, err)
		http.Error(w, fmt.Sprintf("Error while getting rows in %s table ", h.providerTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(providers)
}

func (h *Handlers) CreateSkillHandler(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill
	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	err := h.db.Insert(h.skillTableName, skill)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.skillTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.skillTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(skill)
}

func (h *Handlers) UpdateSkillHandler(w http.ResponseWriter, r *http.Request) {
	var skill models.Skill
	if err := json.NewDecoder(r.Body).Decode(&skill); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	cond := models.SqlCondition{
		Condition: fmt.Sprintf("%s = ?", "skill_id"),
		Values:    []interface{}{skill.SkillID},
	}
	rowsUpdated, err := h.db.Updates(h.skillTableName, skill, cond)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.skillTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.skillTableName), http.StatusInternalServerError)
		return
	}
	if rowsUpdated == 0 {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.skillTableName, fmt.Errorf("No rows updated for condition %v ", cond))
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.skillTableName), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(skill)
}

func (h *Handlers) UpdateTaskProgressHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.taskTableName, fmt.Errorf("No rows updated for condition %v ", cond))
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.taskTableName), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handlers) MakeOfferHandler(w http.ResponseWriter, r *http.Request) {
	var offer models.Offer
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input. Error : %v ", err), http.StatusBadRequest)
		return
	}
	err := h.db.Insert(h.offerTableName, offer)
	if err != nil {
		log.Printf("Error while inserting row in %s table. Error : %v \n", h.offerTableName, err)
		http.Error(w, fmt.Sprintf("Error inserting a row to %s table ", h.offerTableName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(offer)
}
