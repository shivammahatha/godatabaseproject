package migrate

import (
	"github.com/shivammahatha/initializers"
	"github.com/shivammahatha/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDataBase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}