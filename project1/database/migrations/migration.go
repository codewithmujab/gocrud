package migrations

import (
	"fmt"
	"restapi/database"
	"restapi/models"
)

func Migration() {
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		panic("failed to migrate")
	}
	fmt.Println("Migrated successfully")
}
