package contactshandler

import (
	database "github.com/miliyash/ms-contact-manager/app/database"
	dbModels "github.com/miliyash/ms-contact-manager/app/models/database"
)

func GetNewContactHandler() IContact {
	return &Contact{}
}

type Contact dbModels.Contact

type IContact interface {
	CreateContactFromchannel(chContact chan *dbModels.Contact)

	CreateContacts(contacts *[]dbModels.Contact)
	CreateContact(contact *dbModels.Contact)
	UpdateContact()
	DeleteContact()
	GetContact()
}

func (c *Contact) CreateContacts(contacts *[]dbModels.Contact) {

	dbService := database.NewDataBaseService()
	db, err := dbService.GetDB()
	if err == nil {
		if !db.HasTable("contacts") {
			db.CreateTable(&dbModels.Contact{})
		}
		contactsLocal := *contacts
		for _, r := range contactsLocal {
			db.Create(&r)
		}

	}
	db.Close()

}

func (c *Contact) CreateContact(contact *dbModels.Contact) {

	dbService := database.NewDataBaseService()
	db, err := dbService.GetDB()
	if err == nil {
		if !db.HasTable("contacts") {
			db.CreateTable(&dbModels.Contact{})
		}
		db.Create(&contact)
	}
	db.Close()
}

func (c *Contact) CreateContactFromchannel(chContact chan *dbModels.Contact) {

	contactchannel := <-chContact
	dbService := database.NewDataBaseService()
	db, err := dbService.GetDB()
	if err == nil {
		if !db.HasTable("contacts") {
			db.CreateTable(&dbModels.Contact{})
		}
		db.Create(&contactchannel)
	}
	db.Close()
}

func (c *Contact) UpdateContact() {

}

func (c *Contact) DeleteContact() {

}

func (c *Contact) GetContact() {

}
