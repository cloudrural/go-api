package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type rating struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Rate  int64  `json:"rate"`
}

type movie struct {
	Title string `json:"title"`
	Year  int64  `json:"year"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
var ratings = []rating{
	{ID: "1", Title: "Blue Train", Rate: 9},
	{ID: "2", Title: "Jeru", Rate: 9},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Rate: 9},
}

var movies = []movie{
	{Title: "Blue Train", Year: 1993},
	{Title: "Jeru", Year: 1996},
	{Title: "Sarah Vaughan and Clifford Brown", Year: 1991},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
	fmt.Println("Starting the Rest API :::::::::::::)))")
}

func movieList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
	fmt.Println("Starting the Rest API :::::::::::::)))")
}

func getRatings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ratings)
	fmt.Println("Starting the Rest API :::::::::::::)))")
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func postMovies(c *gin.Context) {
	var newMovie movie

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newMovie); err != nil {
		return
	}

	// Add the new album to the slice.
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, newMovie)
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/ratings", getRatings)
	router.GET("/movies", movieList)
	router.POST("/movies", postMovies)
	router.Run("localhost:8080")
}
