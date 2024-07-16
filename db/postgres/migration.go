package postgres

import (
	"fmt"
	"log"
	"time"

	Entity "accessment.com/microservice/db/entity"

	"accessment.com/microservice/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateTables() *gorm.DB {

	// host := utils.GetEnv("POSTGRES_HOST", "")
	user := utils.GetEnv("POSTGRES_USER", "")
	password := utils.GetEnv("POSTGRES_PASSWORD", "")
	dbname := utils.GetEnv("POSTGRES_DB", "")
	port := utils.GetEnv("POSTGRES_PORT", "")

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", user, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Cannot connect to DB at this time, please try again")
	}

	db.AutoMigrate(&Entity.RepoDetail{})
	db.AutoMigrate(&Entity.Commit{})

	postgresDB, err1 := db.DB()
	if err1 == nil {
		postgresDB.SetConnMaxLifetime(time.Millisecond * 200)
	}

	return db

}

func ConnectToDb() *gorm.DB {

	// host := utils.GetEnv("POSTGRES_HOST", "")
	user := utils.GetEnv("POSTGRES_USER", "")
	password := utils.GetEnv("POSTGRES_PASSWORD", "")
	dbname := utils.GetEnv("POSTGRES_DB", "")
	port := utils.GetEnv("POSTGRES_PORT", "")

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", user, password, dbname, port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Cannot connect to DB at this time, please try again")
	}

	postgresDB, err1 := db.DB()
	if err1 == nil {
		postgresDB.SetConnMaxLifetime(time.Millisecond * 200)
	}

	return db
}
