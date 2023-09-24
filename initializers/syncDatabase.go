package initializers

import "goJWT/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
