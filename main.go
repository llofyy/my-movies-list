package main

import (
	"os"
	"project/my-movies-list/movies"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	router.GET("/movies", movies.GetMovies)

	router.Run(":" + os.Getenv("APP_PORT"))
}