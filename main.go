package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
}

var movies = []Movie{
	{ID: "1", Title: "The Shawshank Redemption", Director: "Frank Darabont", Genre: "Drama"},
	{ID: "2", Title: "Oldboy", Director: "Park Chan-wook", Genre: "Mystery/Action"},
	{ID: "3", Title: "2001: A Space Odyssey", Director: "Stanley Kubrick", Genre: "Sci-fi/Adventure"},
}

func getMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

func movieByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := getMovieByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, movie)
}

func getMovieByID(id string) (*Movie, error) {
	for i, m := range movies {
		if m.ID == id {
			return &movies[i], nil
		}
	}

	return nil, errors.New("Movie not found")
}

func createMovie(c *gin.Context) {
	var newMovie Movie

	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func main() {
	router := gin.Default()

	router.GET("/movies", getMovies)
	router.GET("/movies/:id", movieByID)
	router.POST("/movies", createMovie)

	router.Run("localhost:8080")
}
