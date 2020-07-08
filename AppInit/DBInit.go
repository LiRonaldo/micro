package AppInit

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:123456@(localhost:3306)/gomicro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}
func GetDB() *gorm.DB {
	return db
}
