package router

import (
	"github.com/WelintonJunior/SuptecGO/src/server/handlers"
	"github.com/gin-gonic/gin"
)

func UsuRoutes(server *gin.Engine, handler handlers.UsuHandler) {
	server.POST("/login", handler.NewUsu)
}
