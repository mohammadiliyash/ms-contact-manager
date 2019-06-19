package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	dbModels "github.com/miliyash/ms-contact-manager/app/models/database"
	cfg "github.com/miliyash/ms-contact-manager/config"
)

func NewDataBaseService() databaseService {
	return dbStruct{}
}

type databaseService interface {
	GetDB() (db *gorm.DB, err error)
}

type dbStruct struct {
}

var userDB *gorm.DB

func Initialize() {

	var dbStruct = dbStruct{}
	db, err := dbStruct.GetDB()

	if err == nil {
		err = db.DB().Ping()
		if err != nil {
			log.Fatal("DB unreachable:", err)
		}

	}
	db.Close()
}

func (obj dbStruct) GetDB() (db *gorm.DB, err error) {

	conn := fmt.Sprintf("sqlserver://%s:%s@%s:1433?database=contactmanager", cfg.Config.SQLUsername, cfg.Config.SQLPassword, cfg.Config.SQLServer)
	db, err = gorm.Open("mssql", conn)
	if !db.HasTable("Contacts") {
		db.CreateTable(&dbModels.Contact{})
	}
	if !db.HasTable("Customers") {
		db.CreateTable(&dbModels.Customer{})
	}

	return
}
