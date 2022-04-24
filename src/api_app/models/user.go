package models

import (
	"fmt"
	"github.com/api_app/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"ID" gorm:"type:uuid"`
	FirstName string    `json:"FirstName"`
	LastName  string    `json:"LastName"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	IsActive  bool      `json:"IsActive"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type Users []User

func (users *Users) ReadRecords() {
	database.PG.Conn.Find(&users)
}

func (user *User) ReadRecord(searchField string, searchValue any) *gorm.DB {
	query := fmt.Sprintf("%s = ?", searchField)
	tx := database.PG.Conn.First(&user, query, searchValue)
	return tx
}

func (user *User) CreateRecord() *gorm.DB {
	result := database.PG.Conn.Create(&user)
	return result
}

func (user *User) UpdateRecord() *gorm.DB {
	result := database.PG.Conn.Model(&user).Updates(user)
	return result
}
