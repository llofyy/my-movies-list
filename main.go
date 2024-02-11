package main

import (
	"log"
	"os"
	"project/my-movies-list/movies"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var ginMode string;
	environment := os.Getenv("ENVIRONMENT")

	if( environment == "development") {
		ginMode = gin.DebugMode
	} else if environment == "production" {
		ginMode = gin.ReleaseMode
		log.Printf("Server Started")
	}
	
	gin.SetMode(ginMode)

	router := gin.Default()

	router.GET("/movies", movies.GetMovies)

	router.Run(":" + os.Getenv("APP_PORT"))
}