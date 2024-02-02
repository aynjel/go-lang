package migrate

import (
	"go-lang/initializers"
	"go-lang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Post{})
}

func Migrate() {
	initializers.DB.AutoMigrate(&models.Post{})
}
