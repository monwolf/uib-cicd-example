package main

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strings" // para usar función deprecated
)

func main() {
	_ = strings.Title("hello world") // uso de función deprecated para problema de linter

	// Uso del módulo JWT para evitar que go mod tidy lo elimine
	claims := jwt.MapClaims{"user": "test"}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_ = token

	db, _ := sql.Open("sqlite3", ":memory:")
	db.Exec("CREATE TABLE users (id INTEGER, name TEXT)")

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		query := "SELECT * FROM users WHERE name = ?"
		rows, _ := db.Query(query, name)
		defer rows.Close()
		// Simular procesamiento
		var id int
		var userName string
		if rows.Next() {
			rows.Scan(&id, &userName)
		}
		c.JSON(http.StatusOK, gin.H{"id": id, "name": userName})
	})
	r.Run()
}