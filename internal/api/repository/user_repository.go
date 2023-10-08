package repository

import (
	"github.com/abishz17/go-backend-template/internal/domain"
)

type UserRepository struct {
	db *DataBase
}

func NewUserRepository(db *DataBase) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(user *domain.User) (*domain.User, error) {

	err := r.db.conn.Db.Create(&user).Error
	return user, err
}

func (r UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.conn.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
