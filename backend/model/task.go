package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type TaskEntity struct {
	Id          string `json: "id"`
	Name        string `json: "name"`
	Status      int    `json: "status"`
	StatusLabel string `json: "statusLabel"`
	CreatedAt   string `json: "createdAt"`
}

const TASK_STATUS_TODO_ID = 1
const TASK_STATUS_DOING_ID = 2
const TASK_STATUS_DONE_ID = 3

var TaskStatusLabels map[int]string = map[int]string{
	TASK_STATUS_TODO_ID:  "todo",
	TASK_STATUS_DOING_ID: "doing",
	TASK_STATUS_DONE_ID:  "done",
}

type task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *task {
	return &task{db}
}

func (t *task) TaskFind(id int) TaskEntity {
	row := t.db.QueryRow("select id, name, status, created_at from task where id = ?", id)
	var taskEntity TaskEntity
	row.Scan(&taskEntity.Id, &taskEntity.Name, &taskEntity.Status, &taskEntity.CreatedAt)
	taskEntity.StatusLabel = TaskStatusLabels[taskEntity.Status]
	return taskEntity
}

func (t *task) TaskList() []TaskEntity {
	ctx2, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row, err := t.db.QueryContext(ctx2, `select id, name, status, created_at from task order by created_at desc`)
	if err != nil {
		panic(err)
	}
	var taskEntity TaskEntity
	taskList := []TaskEntity{}
	a := []TaskEntity{}
	fmt.Println(a)
	for row.Next() {
		row.Scan(&taskEntity.Id, &taskEntity.Name, &taskEntity.Status, &taskEntity.CreatedAt)
		taskEntity.StatusLabel = TaskStatusLabels[taskEntity.Status]
		taskList = append(taskList, taskEntity)
	}
	return taskList
}

func (t *task) TaskCreate(name string, status int) error {
	_, err := t.db.Exec(`insert into task (name, status, created_at) values (?, ?, now())`, name, status)
	return err
}

func (t *task) TaskUpdate(id int, name string, status int) error {
	_, err := t.db.Exec("update task set name = ?, status = ? where id = ?", name, status, id)
	return err
}
