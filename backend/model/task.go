package model

import (
	"app-runner-study/entity"
	"context"
	"database/sql"
	"time"
)

type TaskEntity struct {
	Id          string `json: "id"`
	Name        string `json: "name"`
	Status      int32  `json: "status"`
	StatusLabel string `json: "statusLabel"`
	CreatedAt   string `json: "createdAt"`
}

const TASK_STATUS_TODO_ID = 1
const TASK_STATUS_DOING_ID = 2
const TASK_STATUS_DONE_ID = 3

var TaskStatusLabels map[int32]string = map[int32]string{
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
	query := entity.New(t.db)
	tasks, err := query.LisTasks(ctx2)
	if err != nil {
		panic(err)
	}
	var taskEntity TaskEntity
	taskList := []TaskEntity{}
	for _, t := range tasks {
		taskEntity.Name = t.Name
		taskEntity.CreatedAt = t.CreatedAt.String()
		taskEntity.Status = t.Status
		taskEntity.StatusLabel = TaskStatusLabels[t.Status]
		taskList = append(taskList, taskEntity)
	}
	return taskList
}

func (t *task) TaskCreate(name string, status int32) error {
	query := entity.New(t.db)

	_, err := query.CreateTask(context.Background(), entity.CreateTaskParams{
		Name:   name,
		Status: status,
	})
	return err
}

func (t *task) TaskUpdate(id int64, name string, status int) error {
	query := entity.New(t.db)
	query.UpdateTask(context.Background(), entity.UpdateTaskParams{
		Name:   name,
		Status: TASK_STATUS_TODO_ID,
		ID:     id,
	})
	_, err := t.db.Exec("update task set name = ?, status = ? where id = ?", name, status, id)
	return err
}
