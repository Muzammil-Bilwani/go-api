package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080";
	// Create a new Gin router
	api := gin.Default()

	api.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Api is working",
		})
	})

	api.GET("/locations", func(c *gin.Context) {
		cities := []string{"New York", "London", "Tokyo", "Paris"}
		c.JSON(200, gin.H{"data": cities})
	})

	api.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")
		weather := generateRandomWeather(city)
		c.JSON(200, weather)
	})

	api.Run(port)
}

// generateRandomWeather generates a random weather response for the given city
func generateRandomWeather(city string) gin.H {
	rand.Seed(time.Now().UnixNano())

	// Weather conditions and descriptions
	conditions := []string{"Sunny", "Cloudy", "Rainy", "Snowy"}
	descriptions := []string{"Clear skies", "Partly cloudy", "Chance of showers", "Heavy snowfall"}

	// Generate random weather data
	temperature := rand.Intn(40) // Random temperature between 0 and 40
	humidity := rand.Intn(100)   // Random humidity between 0 and 100
	condition := conditions[rand.Intn(len(conditions))]         // Random condition from the conditions list
	description := descriptions[rand.Intn(len(descriptions))]   // Random description from the descriptions list

	// Create the weather response
	weather := gin.H{
		"city":        city,
		"temperature": temperature,
		"humidity":    humidity,
		"condition":   condition,
		"description": description,
	}

	return weather
}
