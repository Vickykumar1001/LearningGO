package repository

import (
	"context"
	"errors"
	"log"
	"session-23-gin-jwt/internal/models"

	"gorm.io/gorm"
)

type MysqlOrm struct {
	db *gorm.DB
}

func (m MysqlOrm) CreateUser(ctx context.Context, user models.User) (interface{}, error) {
	tx := m.db.Create(&user)
	if tx.Error != nil {
		log.Println("error", tx.Error)
		return 0, tx.Error
	}

	return user.ID, nil
}

func (m MysqlOrm) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	log.Println("Incoming username", userName)
	tx := m.db.Where("UserName = ?", userName).Find(&user)
	if tx.Error != nil {
		log.Println("error", tx.Error)
		return nil, tx.Error
	}
	if user.Username == "" {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

func (m MysqlOrm) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	// var user models.User
	var users []*models.User
	result := m.db.Find(&users)
	return users, result.Error
}

func (m MysqlOrm) UpdateUser(ctx context.Context, id interface{}, user models.User) error {
	// fmt.Println(user)
	m.db.Save(&user)
	return nil
}

func (m MysqlOrm) DeleteUser(ctx context.Context, ID interface{}) error {
	m.db.Where("ID = ?", ID).Delete(&models.User{})
	return nil
}

func NewMysqlOrm(db *gorm.DB) DbRepository {
	return &MysqlOrm{db: db}
}
