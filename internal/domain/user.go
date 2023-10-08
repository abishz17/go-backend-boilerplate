package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID      `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	IsAdmin   bool           `json:"is_admin"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	id := uuid.New()
	u.ID = id
	u.IsAdmin = true
	return nil
}
