package initializers

import "github.com/ajujacob88/go-jwt-auth-basic-webapp/models"

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}
