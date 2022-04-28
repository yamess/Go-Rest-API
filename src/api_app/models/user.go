package models

import (
	"encoding/json"
	"fmt"
	"github.com/api_app/database"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"email"`
	Password  string    `json:"password" validate:"required"`
	IsActive  bool      `json:"is_active" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt NullTime  `json:"updated_at" gorm:"default:null"`
}

type Users []User

func (user *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(user)
}

func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}

func (users *Users) ReadRecords() {
	database.PG.Conn.Find(&users)
}

func (user *User) ReadRecord(searchField string, searchValue any) *gorm.DB {
	query := fmt.Sprintf("%s = ?", searchField)
	result := database.PG.Conn.First(&user, query, searchValue)
	return result
}

func (user *User) CreateRecord() *gorm.DB {
	result := database.PG.Conn.Create(&user)
	return result
}

func (user *User) UpdateRecord(id string) *gorm.DB {
	user.ID = uuid.Must(uuid.Parse(id))

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

// BeforeCreate Hook
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()
	return nil
}
