package db

import (
	"log"
	interfaces "marketplace/internal/interfaces/database"
	"marketplace/internal/models"

	"gorm.io/gorm"
)

type dbData struct {
	db *gorm.DB
}

func NewDb(userName, password, host, dbName string, port int, providerTableName, userTableName, skillTableName, taskTableName, offerTableName string) (interfaces.DbOps, error) {
	gormDb, err := connectToDb(userName, password, host, dbName, port)

	err = gormDb.Table(providerTableName).AutoMigrate(&models.Provider{})
	if err != nil {
		log.Printf("Error while automigating %s table . Error : %v \n", providerTableName, err)
		return nil, err
	}

	err = gormDb.Table(userTableName).AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("Error while automigating %s table . Error : %v \n", userTableName, err)
		return nil, err
	}

	err = gormDb.Table(skillTableName).AutoMigrate(&models.Skill{})
	if err != nil {
		log.Printf("Error while automigating %s table . Error : %v \n", skillTableName, err)
		return nil, err
	}

	err = gormDb.Table(taskTableName).AutoMigrate(&models.Task{})
	if err != nil {
		log.Printf("Error while automigating %s table . Error : %v \n", taskTableName, err)
		return nil, err
	}

	err = gormDb.Table(offerTableName).AutoMigrate(&models.Offer{})
	if err != nil {
		log.Printf("Error while automigating %s table . Error : %v \n", offerTableName, err)
		return nil, err
	}

	return &dbData{db: gormDb}, err
}

func (d *dbData) Get(tableName string, data interface{}, condition models.SqlCondition, limit int, offset int, order string) error {
	log.Printf("Querying %s table with condition %#v, for limit %d offset %d and order %s \n", tableName, condition, limit, offset, order)
	output := d.db.Table(tableName).Where(condition.Condition, condition.Values...).Limit(limit).Offset(offset).Order(order).Find(data)
	if output.Error != nil {
		log.Printf("Error while getting data from table %s . Error : %v \n", tableName, output.Error)
		return output.Error
	}
	log.Printf("%d rows got while querying to table %s \n", output.RowsAffected, tableName)
	return nil
}

func (w *dbData) Insert(tableName string, data interface{}) error {
	output := w.db.Table(tableName).Create(data)
	if output.Error != nil {
		log.Printf("Error while inserting row to table %s. Error : %v \n", tableName, output.Error)
		return output.Error
	}
	log.Printf("%d rows affected while inserting to table %s \n", output.RowsAffected, tableName)
	return nil
}

func (w *dbData) Updates(tableName string, data interface{}, condition models.SqlCondition) (int, error) {
	log.Printf("Updating table %s rows for condition %#v with value %#v \n", tableName, condition, data)
	output := w.db.Table(tableName).Where(condition.Condition, condition.Values...).Updates(data)
	if output.Error != nil {
		log.Printf("Error while updating row to table %s, with condition %#v . Error : %v \n", tableName, condition, output.Error)
		return 0, output.Error
	}
	log.Printf("%d rows affected while updating table %s \n", output.RowsAffected, tableName)
	return int(output.RowsAffected), nil
}

func (w *dbData) Delete(tableName string, condition models.SqlCondition) (int, error) {
	log.Printf("Deleting table %s rows for condition %#v \n", tableName, condition)
	output := w.db.Table(tableName).Where(condition.Condition, condition.Values...).Delete(nil)
	if output.Error != nil {
		log.Printf("Error while delting rows to table %s, with condition %#v . Error : %v \n", tableName, condition, output.Error)
		return 0, output.Error
	}
	log.Printf("%d rows affected while deleting table %s \n", output.RowsAffected, tableName)
	return int(output.RowsAffected), nil
}
