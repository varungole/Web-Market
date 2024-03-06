package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var products = []product{}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", addProducts)
	router.GET("/products/:id", getProductsById)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func addProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductsById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of products, looking for
	// a product whose ID value matches the parameter.
	for _, p := range products {
		if strings.ToLower(p.ID) == strings.ToLower(id) {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
