package contactshandler

import (
	database "github.com/miliyash/ms-contact-manager/app/database"
	dbModels "github.com/miliyash/ms-contact-manager/app/models/database"
)

func GetNewCustomerHandler() ICustomer {
	return &customer{}
}

type customer dbModels.Customer

type ICustomer interface {
	CreateCustomer()
	UpdateCustomer()
	DeleteCustomer()
	GetCustomer()
}

func (c *customer) CreateCustomer() {

	dbService := database.NewDataBaseService()
	db, err := dbService.GetDB()
	if err != nil {
		if !db.HasTable("contacts") {
			db.Debug().AutoMigrate(&dbModels.Customer{})
		} else {
			db.Create(&c)
		}

	}

}

func (c *customer) UpdateCustomer() {

}

func (c *customer) DeleteCustomer() {

}

func (c *customer) GetCustomer() {

}
