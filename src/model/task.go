package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type TaskEntity struct {
	Name      string `json: "name"`
	CreatedAt string `json: "createdAt"`
}

type task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *task {
	return &task{db}
}

func (t *task) TaskList() []TaskEntity {
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
	return taskList
}
func (t *task) TaskCreate(name string) {
	t.db.Exec(`insert into task (name, created_at) values (?, now())`, name)
}
