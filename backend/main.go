package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var db *sql.DB

func main() {

	dsn := "root:tereocta12@tcp(127.0.0.1:3306)/inventory_db?parseTime=true"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	router.GET("/products", func(c *gin.Context) {

		var products []Product

		rows, err := db.Query("SELECT id, name, price, quantity FROM products")
		if err != nil {
			log.Fatal("Error querying the database: ", err)
		}
		defer rows.Close()

		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
				log.Fatal("Error scanning row: ", err)
			}
			products = append(products, p)
		}

		c.IndentedJSON(http.StatusOK, products)
	})

	router.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.BindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)"
		result, err := db.Exec(query, newProduct.Name, newProduct.Price, newProduct.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		id, _ := result.LastInsertId()
		newProduct.ID = int(id)

		c.IndentedJSON(http.StatusCreated, newProduct)
	})

	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product

		if err := c.BindJSON(&updatedProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID JSON"})
			return
		}

		query := "UPDATE products SET name = ?, price = ?, quantity = ? WHERE id = ?"
		_, err := db.Exec(query, updatedProduct.Name, updatedProduct.Price, updatedProduct.Quantity, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, updatedProduct)
	})

	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		query := "DELETE FROM products WHERE id = ?"

		_, err := db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	})

	router.Run("localhost:8080")
}
