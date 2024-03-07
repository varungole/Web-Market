package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var db *sql.DB

func main() {
	fmt.Println("Welcome to DB Connection")

	// Initialize the MySQL driver
	var err error
	db, err = sql.Open("mysql", ":@A@tcp(localhost:3306)/webMarket")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Ping the database to ensure a connection can be established
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to DB")

	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", addProducts)

	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM PRODUCTS")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []product

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Title, &p.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func addProducts(c *gin.Context) {
	var p product

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO products (id, title, price) VALUES (?, ?, ?)", p.ID, p.Title, p.Price)
	if err != nil {
		fmt.Println("Error executing SQL query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}
