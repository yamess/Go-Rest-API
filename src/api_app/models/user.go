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
	UpdatedAt time.Time `json:"UpdatedAt"`
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
	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	result := database.PG.Conn.Create(&user)
	return result
}

func (user *User) UpdateRecord(id string) *gorm.DB {
	user.ID = uuid.Must(uuid.Parse(id))
	user.UpdatedAt = time.Now().UTC()

	result := database.PG.Conn.Model(&user).
		Omit("id", "created_at").
		Updates(user)
	if result.Error == nil {
		user.ReadRecord("id", user.ID)
		user.Password = ""
	}

	return result
}

func (user *User) DeleteRecord(id string) *gorm.DB {
	user.ID = uuid.Must(uuid.Parse(id))
	result := database.PG.Conn.Delete(&user)
	return result
}
