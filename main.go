package main

import (
	"project/my-movies-list/movies"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/movies", movies.GetMovies)
	router.POST("/movies", movies.AddMovie)

	router.Run("localhost:8080")
}