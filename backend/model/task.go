package model

import (
	"context"
	"database/sql"
	"fmt"
	"task-app-study/entity"
	"time"
)

type TaskEntity struct {
	Id          int64  `json: "id"`
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

func (t *task) TaskFind(id int64) TaskEntity {
	query := entity.New(t.db)
	task, err := query.GetTask(context.Background(), id)
	if err != nil {
		return TaskEntity{}
	}
	fmt.Println(task)
	// row := t.db.QueryRow("select id, name, status, created_at from task where id = ?", id)
	var taskEntity TaskEntity = TaskEntity{
		Id:          task.ID,
		Name:        task.Name,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.String(),
		StatusLabel: TaskStatusLabels[task.Status],
	}
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
		taskEntity.Id = t.ID
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
	return query.UpdateTask(context.Background(), entity.UpdateTaskParams{
		Name:   name,
		Status: TASK_STATUS_TODO_ID,
		ID:     id,
	})
}

func (t *task) TaskDelete(id int64) error {
	query := entity.New(t.db)
	err := query.DeleteTask(context.Background(), id)
	return err
}
