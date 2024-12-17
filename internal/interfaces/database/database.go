package interfaces

import "marketplace/internal/models"

type DbOps interface {
	Insert(string, interface{}) error
	Updates(string, interface{}, models.SqlCondition) (int, error)
	Delete(string, models.SqlCondition) (int, error)
	Get(string, interface{}, models.SqlCondition, int, int, string) error
}
