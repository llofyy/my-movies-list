package movies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieType struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Stars int    `json:"stars"`
	Image string `json:"image"`
}

var movies = []MovieType{
	{Id:    "akluhsiahsihaishaisha",
	Name:  "Good Burger",
	Stars: 3,
	Image: "https://m.media-amazon.com/images/M/MV5BMGM1YTJhOTQtNmYwZS00MGU0LTkxNWEtZGEwM2I4NTc5MzZlXkEyXkFqcGdeQXVyMTUyOTc1NDYz._V1_FMjpg_UX1000_.jpg",},
}

func AddMovie(c *gin.Context) {
	var newMovie MovieType;

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, movies)
}

func GetMovies(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, movies)
}