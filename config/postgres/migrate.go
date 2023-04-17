package postgres

import (
	"log"

	"github.com/gvriofernando/test_saham_rakyat/domain"
	"gorm.io/gorm"
)

// DBMigrate will create & migrate the tables, then make the some relationships if necessary
func DBMigrate(conn *gorm.DB) error {
	if conn.Error != nil {
		return conn.Error
	}

	conn.AutoMigrate(domain.User{})
	log.Println("Migration has been processed")

	return nil
}
