package dboperation

import "github.com/jinzhu/gorm"

const (
	DB_HOST = "localhost"
	DB_PORT = "33062"
	DB_USER = "root"
	DB_PASS = "my-secret-pw"
	DB_NAME = "webpro-db"
)

func gormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(localhost:33062)/webpro-db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
