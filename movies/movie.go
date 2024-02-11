package movies

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type MovieType struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Stars int    `json:"stars"`
	Image string `json:"image"`
}

func GetMovies(c *gin.Context) {
	URL := os.Getenv("MOVIE_URL")
	
	req, _ := http.NewRequest("GET", URL, nil)
	
	req.Header.Set("Authorization", "Bearer "+ os.Getenv("MOVIE_API_TOKEN"))
	req.Header.Set("accept", "application/json")
	
	movies, err := http.DefaultClient.Do(req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "System Error")
	}
	
	body, _ := io.ReadAll(movies.Body)

	var resp map[string]interface{}
	json.Unmarshal([]byte(body), &resp)

	c.JSON(http.StatusOK, resp)
}