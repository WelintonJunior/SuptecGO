package handlers

import (
	"fmt"
	"net/http"

	"github.com/WelintonJunior/SuptecGO/src/server/adapters/domain"
	"github.com/WelintonJunior/SuptecGO/src/server/app/application"
	"github.com/gin-gonic/gin"
)

type UsuHandler struct {
	service application.UsuService
}

func NewUsuHandler(service application.UsuService) *UsuHandler {
	return &UsuHandler{service: service}
}

func (h *UsuHandler) NewUsu(c *gin.Context) {
	var usu domain.USU_users
	if err := c.ShouldBindJSON(&usu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao receber dados para criação do usuário", "error": true})
		return
	}

	if err := h.service.NewUsu(usu); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro na criação do usuário", "error": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucesso", "error": false})

}
