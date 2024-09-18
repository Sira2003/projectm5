package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/projectm4/config"
	"example.com/projectm4/entity"
)

// GET /genders
func ListBloodTypes(c *gin.Context) {
	var bloodTypes []entity.BloodType
	db := config.DB()
	db.Find(&bloodTypes)
	c.JSON(http.StatusOK, &bloodTypes)
}