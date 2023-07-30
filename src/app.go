package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/task", func(ctx *gin.Context) {
		mysqlUser := os.Getenv("MYSQL_USER")
		mysqlPassword := os.Getenv("MYSQL_PASSWORD")
		mysqlHost := os.Getenv("MYSQL_HOST")
		dsn := fmt.Sprintf("%s:%s@(%s:3306)/information_schema", mysqlUser, mysqlPassword, mysqlHost)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		row, err := db.QueryContext(ctx2, `select table_name from tables`)
		if err != nil {
			panic(err)
		}
		row.Next()
		var a string
		row.Scan(&a)

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Title",
			"count": a,
		})
	})
	r.Run()
}
