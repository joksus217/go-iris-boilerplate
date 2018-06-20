package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "db_go_blog"
)

var db_info = fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_NAME)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", db_info)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	db.SingularTable(true)
	DB = db
	return DB
}

func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("mysql", db_info)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	test_db.LogMode(true)
	DB = test_db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
