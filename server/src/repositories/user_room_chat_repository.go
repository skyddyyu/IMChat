package repositories

import (
	"ImChat/src/models"

	"gorm.io/gorm"
)

type UserRoomChatRepository interface {
	JoinChatRoom(record *models.UserChatRoom) error
	FindChatRoomUser(record *models.UserChatRoom) error
}

func NewUserRoomChatRepository(db *gorm.DB) UserRoomChatRepository {
	return &userRoomChatRepository{db}
}

type userRoomChatRepository struct {
	db *gorm.DB
}

func (r *userRoomChatRepository) JoinChatRoom(record *models.UserChatRoom) error {
	return r.db.Create(record).Error
}

func (r *userRoomChatRepository) FindChatRoomUser(record *models.UserChatRoom) error {
	return r.db.Where(record).First(record).Error
}