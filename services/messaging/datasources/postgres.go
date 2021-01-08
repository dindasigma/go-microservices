package datasources

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

func InitializePostgres(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}
}
