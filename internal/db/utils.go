package db

import (
	"fmt"
	"log"

	mysql_driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectToDb(userName, password, host, dbName string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, host, port, dbName)
	dsnConf, err := mysql_driver.ParseDSN(dsn)
	if err != nil {
		log.Printf("Error while parsing DSN. Error : %v ", err)
		return nil, err
	}
	mysqlDialector := &mysql.Dialector{Config: &mysql.Config{DSN: dsn, DSNConfig: dsnConf}}
	return gorm.Open(mysqlDialector)
}
