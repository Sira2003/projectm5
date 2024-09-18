package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/projectm4/config"
	"example.com/projectm4/entity"
)

// GET /jobPositions
func ListJobPositions(c *gin.Context) {
	var jobPositions []entity.JobPosition
	db := config.DB()
	db.Find(&jobPositions)
	c.JSON(http.StatusOK, &jobPositions)
}
