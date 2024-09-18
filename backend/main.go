package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"example.com/projectm4/config"
	"example.com/projectm4/controller"
)

const PORT = "8000"

func main() {
	// open connection database
	config.ConnectionDB()
	// Generate databases
	config.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	router := r.Group("")
	{
		// Employee Routes
		router.GET("/employees", controller.ListEmployees)
		router.GET("/employee/:id", controller.GetEmployee)
		router.POST("/employees", controller.CreateEmployee)
		router.PATCH("/employees", controller.UpdateEmployee)
		router.DELETE("/employees/:id", controller.DeleteEmployee)
		// Gender Routes
		router.GET("/genders", controller.ListGenders)
		// BloodType Routes
		router.GET("/bloodTypes", controller.ListBloodTypes)
		// JobPosition Routes
		router.GET("/jobPositions",controller.ListJobPositions)
		// Patient Routes
		router.GET("/patients", controller.ListPatients)
		router.GET("/patient/:id", controller.GetPatient)
		router.POST("/patients", controller.CreatePatient)
		router.PATCH("/patients", controller.UpdatePatient)
		router.DELETE("/patients/:id", controller.DeletePatient)

		//

		//
	}
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// Run the server

	r.Run("localhost:" + PORT)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
