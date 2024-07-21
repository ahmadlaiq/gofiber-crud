package migrations

import (
	"fmt"
	"log"

	"fiber-rest-api/database"
	"fiber-rest-api/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal("Failed to migrate...")
	}

	fmt.Println("Migrated successfully")
}
