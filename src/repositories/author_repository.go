package repositories

import (
	"models"
	"common"
	"github.com/jinzhu/gorm"
	"fmt"
)

var authors []models.Author
var author models.Author

type AuthorRepository interface {
	GetAll() (result *gorm.DB)
	GetByID(id int64) (result *gorm.DB)
}

func GetAll() (result *gorm.DB, err error) {
	result = common.GetDB().Find(&authors)
	if err = result.Error; err != nil {
		fmt.Printf("AuthorRepository GetAll Error: %v", err)
	}
	return result, err
}

func GetByID(id int64) (result *gorm.DB, err error) {
	result = common.GetDB().Find(&authors, id)
	if err = result.Error; err != nil {
		fmt.Printf("AuthorRepository GetAll Error: %v", err)
	}
	return result, err
}
