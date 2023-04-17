package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	User     string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func Init(cfg ConfigDB) *gorm.DB {
	var err error
	var db *gorm.DB
	user := cfg.User
	password := cfg.Password
	host := cfg.Host
	port := cfg.Port
	database := cfg.Dbname

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}
	log.Println("Database connected")

	DBMigrate(db)
	return db
}
