/*
 * @File: controllers.user.go
 * @Description: Implementação API de Testes
 * @Author: Carlos Henrique Lemos (chenriquelemos@gmail.com)
 */
package controllers

import (
	"net/http"

	"../models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Custom struct {
	Resultado string
}

// Test godoc
// @Summary Checa se o serviço está funcinando
// @Description Api de Teste com Golang + Gin Web(https://github.com/gin-gonic/gin)
// @Tags test
// @Produce  json
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Success 200
// @Router /test/hello [get]
func (c *Custom) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Message{"Successfully"})
	log.Debug("[INFO]: ", "Chamou o serviço Test dentro do Custom!")
}
