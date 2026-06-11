package main

import (
	"github.com/testdev-learning/DataHub/backend/config"
	"github.com/testdev-learning/DataHub/backend/routes"
	"github.com/testdev-learning/DataHub/backend/utils"
)

func main() {
	config.InitConfig()
	utils.InitDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}