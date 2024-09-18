package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/projectm4/config"
	"example.com/projectm4/entity"
)

// GET /genders
func ListGenders(c *gin.Context) {
	var genders []entity.Gender
	db := config.DB()
	db.Find(&genders)
	c.JSON(http.StatusOK, &genders)
}
