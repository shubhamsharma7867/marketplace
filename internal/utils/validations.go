package utils

import (
	"marketplace/internal/models"

	validator "github.com/go-playground/validator/v10"
)

type Validator struct {
}

func (v *Validator) Validate(companyfeilds models.CompanyFields) bool {
	err := validator.New().Struct(companyfeilds)
	return err == nil
}
