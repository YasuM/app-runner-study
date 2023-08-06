package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"

	"app-runner-study/handler"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))
	r.LoadHTMLGlob("templates/*")
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/task", mysqlUser, mysqlPassword, mysqlHost)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	thandler := handler.NewTaskHandler(db)
	r.GET("/task", thandler.TaskList)
	r.GET("/api/task", thandler.TaskListApi)
	r.GET("/input", thandler.TaskInput)
	r.POST("/create", thandler.TaskCreate)
	r.Run()
}
