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

type task struct {
	db *sql.DB
}

type TaskEntity struct {
	Name      string
	CreatedAt string
}

func (t *task) taskList(ctx *gin.Context) {
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row, err := t.db.QueryContext(ctx2, `select name, created_at from task`)
	if err != nil {
		panic(err)
	}
	var taskEntity TaskEntity
	taskList := []TaskEntity{}
	a := []TaskEntity{}
	fmt.Println(a)
	for row.Next() {
		row.Scan(&taskEntity.Name, &taskEntity.CreatedAt)
		taskList = append(taskList, taskEntity)
	}

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "Task",
		"taskList": taskList,
	})
}

func main() {
	r := gin.Default()
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
	t := &task{
		db: db,
	}
	r.GET("/task", t.taskList)
	r.Run()
}
