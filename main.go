package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

var quotes = []string{
	"It's just a flesh wound.",
	"We are the knights who say Ni!",
	"What is the air-speed velocity of an unladen swallow?",
	"I fart in your general direction.",
	"Help! Help! I'm being repressed!",
}

func main() {
	// Create a default Gin router
	router := gin.Default()

	// GET /quotes - return all Monty Python quotes
	router.GET("/quotes", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"quotes": quotes})
	})

	// GET /quote - return a random Monty Python quote
	router.GET("/quote", func(c *gin.Context) {
		index := rand.Intn(len(quotes))
		c.JSON(http.StatusOK, gin.H{"quote": quotes[index]})
	})

	// POST /quote - add a new Monty Python quote to the list
	router.POST("/quote", func(c *gin.Context) {
		var newQuote struct {
			Quote string `json:"quote" binding:"required"`
		}
		if err := c.ShouldBindJSON(&newQuote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		quotes = append(quotes, newQuote.Quote)
		c.JSON(http.StatusOK, gin.H{"message": "Quote added successfully!"})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
