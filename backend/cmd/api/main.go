package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/redis/go-redis/v9"

	"task-app-study/handler"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setCors(r)
	redis := newRedisClient()
	db := newMySQLClient()
	defer db.Close()
	thandler := handler.NewTaskHandler(db, redis)
	uhandler := handler.NewUserHandler(db)
	lhandler := handler.NewLoginHandler(db, redis)
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.GET("/task/:id", thandler.Task)
	r.GET("/task", thandler.TaskList)
	r.GET("/task_status", thandler.TaskStatusList)
	r.POST("/task/create", thandler.TaskCreate)
	r.POST("/task/edit", thandler.TaskEdit)
	r.POST("/task/delete/:id", thandler.TaskDelete)
	r.POST("/user/create", uhandler.UserCreate)
	r.POST("/login", lhandler.Login)
	r.RunTLS(":8080", "/app/cmd/api/localhost.pem", "/app/cmd/api/localhost-key.pem")
}

func setCors(r *gin.Engine) {
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
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
}

func newMySQLClient() *sql.DB {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/task?parseTime=true", mysqlUser, mysqlPassword, mysqlHost)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func newRedisClient() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	if redis == nil {
		panic(redis)
	}
	return redis
}
