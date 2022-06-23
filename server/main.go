package main

import (
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/Anirudh-rao/CalorieTracker-Go-React/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	//Gin.logger will provide us with a Middleware to monitor the Website
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.Getentries)
	router.GET("/entry/:id/", routes.GetentryById)
	router.GET("/ingredient/:ingredient", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id",routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id",  routes.DeleteEntry)

	router.Run(":"+port)
}
