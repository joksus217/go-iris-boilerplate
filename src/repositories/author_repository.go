package repositories

import (
	"models"
	"common"
	"github.com/jinzhu/gorm"
)

type AuthorRepository interface {
	GetAll() (result *gorm.DB)
	GetByID(id int64) (result *gorm.DB)
}

func GetAll() (result *gorm.DB) {
	var authors []models.Author
	return common.GetDB().Find(&authors)
}

func GetByID(id int64) (result *gorm.DB) {
	var author models.Author
	return common.GetDB().First(&author, id)
}

func Insert(author models.Author) (result *gorm.DB) {
	tx := common.GetDB().Begin()

	if err := tx.Create(&author).Error; err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}
