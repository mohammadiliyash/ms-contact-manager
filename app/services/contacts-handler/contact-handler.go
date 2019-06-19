package contactshandler

import (
	database "github.com/miliyash/ms-contact-manager/app/database"
	dbModels "github.com/miliyash/ms-contact-manager/app/models/database"
)

//GetNewContactHandler ...
func GetNewContactHandler() IContact {
	return &Contact{}
}

// Contact ...
type Contact dbModels.Contact

//IContact ...
type IContact interface {
	CreateContactFromchannel(chContact chan *dbModels.Contact)

	CreateContacts(contacts *[]dbModels.Contact)
	CreateContact(contact *dbModels.Contact)
	UpdateContact()
	DeleteContact()
	GetContact()
}

//CreateContacts ...
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

// CreateContact ...
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

// CreateContactFromchannel ...
func (c *Contact) CreateContactFromchannel(chContact chan *dbModels.Contact) {

	contactchannel := <-chContact
	go c.CreateContact(contactchannel)
}

// UpdateContact ...
func (c *Contact) UpdateContact() {

}

// DeleteContact ...
func (c *Contact) DeleteContact() {

}

// GetContact ...
func (c *Contact) GetContact() {

}
