package main

import (
	"github.com/WelintonJunior/SuptecGO/src/server/adapters/repository"
	db "github.com/WelintonJunior/SuptecGO/src/server/adapters/repository/database"
	"github.com/WelintonJunior/SuptecGO/src/server/app/application"
	"github.com/WelintonJunior/SuptecGO/src/server/handlers"
	"github.com/WelintonJunior/SuptecGO/src/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	usuRepo := repository.NewLocalUsuRepository()
	usuService := application.NewUsuService(usuRepo)
	usuHandler := handlers.NewUsuHandler(*usuService)
	router.UsuRoutes(server, *usuHandler)

	server.Run(":3000")
}
