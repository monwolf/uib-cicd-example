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

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE users (id INTEGER, name TEXT)")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		query := "SELECT * FROM users WHERE name = '" + name + "'" // SQL injection vulnerable
		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		defer func() {
			err = rows.Close()
			if err != nil {
				panic(err)
			}
		}()
		// Simular procesamiento
		var id int
		var userName string
		if rows.Next() {
			err = rows.Scan(&id, &userName)
			if err != nil {
				panic(err)
			}
		}
		c.JSON(http.StatusOK, gin.H{"id": id, "name": userName})
	})
	err = r.Run()
	if err != nil {
		panic(err)
	}
}