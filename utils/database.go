package utils

import (
	"fmt"
	"log"

	"github.com/devmeireles/nextjob/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	db, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}

	db.AutoMigrate(
		models.Skill{},
		models.User{},
		models.Address{},
	)

	db.Model(models.Address{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

}

func DBConn() *gorm.DB {
	return db
}
