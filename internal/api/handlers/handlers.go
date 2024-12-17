package handlers

import (
	interfaces "marketplace/internal/interfaces/database"
	"marketplace/internal/utils"
)

type Handlers struct {
	db                interfaces.DbOps
	providerTableName string
	skillTableName    string
	taskTableName     string
	offerTableName    string
	userTableName     string
	requestValidator  utils.Validator
}

func NewHandlers(db interfaces.DbOps, providerTableName string, skillTableName string,
	taskTableName string, offerTableName string, userTableName string, requestValidator utils.Validator) Handlers {
	return Handlers{
		db:                db,
		providerTableName: providerTableName,
		skillTableName:    skillTableName,
		taskTableName:     taskTableName,
		offerTableName:    offerTableName,
		userTableName:     userTableName,
		requestValidator:  requestValidator,
	}
}
