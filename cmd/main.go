package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"./config"
	"./handlers/dashboards"
)

func main() {
	router := gin.Default()

	// Connect to database
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// Set up route handlers
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	dashboard1Handler := dashboards.NewDashboard1Handler(db)
	router.GET("/dashboard1", dashboard1Handler.HandleDashboard1)
	// ... Add routes for other dashboards ...

	// Serve the application
	router.Run(":8080") // Adjust port number as needed
}
