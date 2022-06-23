package main

import (
	"log"

	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type gadget struct {
	ID     string `json:"id"`
	Brand  string `json:"title"`
	Gadget string `json:"artist"`
	Price  int    `json:"price"`
}

var gadgets = []gadget{
	{ID: "1", Brand: "HP", Gadget: "Laptops", Price: 56000},
	{ID: "2", Brand: "MI", Gadget: "Smart Watch", Price: 3000},
	{ID: "3", Brand: "DELL", Gadget: "Laptops", Price: 70000},
	{ID: "4", Brand: "HAVELLS", Gadget: "Electric Motor", Price: 3200},
}

func getCorsConfig() gin.HandlerFunc {
	var origins = []string{"*"}

	return cors.New(cors.Config{

		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Accept", "Accept-Encoding",
			"Accept-Language", "Connection", "Content-Length",
			"Content-Type", "Host", "Origin", "Referer", "User-Agent", "transformRequest"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func main() {
	router := gin.Default()
	router.Use(getCorsConfig())
	router.GET("/gadgets", getgadgets)
	router.GET("/gadgets/:id", getGadgetByID)
	router.POST("/gadgets", postgadgets)
	go func() {
		// service connections
		if err := router.Run(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	router.Run("localhost:8088")
}

func getgadgets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gadgets)
}

func postgadgets(c *gin.Context) {
	var newGadget gadget
	if err := c.BindJSON(&newGadget); err != nil {
		return
	}
	gadgets = append(gadgets, newGadget)
	c.IndentedJSON(http.StatusCreated, newGadget)
}

func getGadgetByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range gadgets {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Gadget not found"})
}
