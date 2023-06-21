package repository

import (
	"context"

	"github.com/ifty123/simple_online_store/internal/dto"
	"github.com/ifty123/simple_online_store/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(ctx context.Context, id uint) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)
	SaveUser(ctx context.Context, user *dto.RegisterUserReq) (*model.User, error)
}

type user struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{
		Db: db,
	}
}

func (e *user) FindById(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	q := e.Db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id)

	err := q.First(&user).Error
	return &user, err
}

func (e *user) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := e.Db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (e *user) SaveUser(ctx context.Context, user *dto.RegisterUserReq) (*model.User, error) {
	newUser := model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := e.Db.WithContext(ctx).Save(&newUser).Error; err != nil {
		return &newUser, err
	}

	return &newUser, nil
}

func (e *user) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var (
		count   int64
		isExist bool
	)

	if err := e.Db.WithContext(ctx).Model(&model.User{}).Where("email = ?", &email).Count(&count).Error; err != nil {
		return isExist, nil
	}

	if count > 0 {
		isExist = true
	}

	return isExist, nil
}
