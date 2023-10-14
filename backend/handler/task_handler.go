package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"task-app-study/model"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type taskHandler struct {
	db    *sql.DB
	redis *redis.Client
}

type TaskForm struct {
	Name string `json:"task" binding:"required"`
}

type TaskFormEditView struct {
	Id int64 `uri:"id" binding:"required"`
}

type TaskFormEdit struct {
	Id     int64  `json:"id" binding:"required"`
	Name   string `json:"task" binding:"required"`
	Status int    `json:"status" binding:"required"`
}

type TaskFormDelete struct {
	Id int64 `uri:"id" binding:"required"`
}

func NewTaskHandler(db *sql.DB, redis *redis.Client) *taskHandler {
	return &taskHandler{db, redis}
}

func (t *taskHandler) Task(ctx *gin.Context) {
	var form TaskFormEditView
	if err := ctx.ShouldBindUri(&form); err != nil {
		fmt.Println(ctx.Param("id"))
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	task := model.NewTask(t.db)
	task.TaskFind(form.Id)
	ctx.JSON(http.StatusOK, task.TaskFind(form.Id))
}

func (t *taskHandler) TaskList(ctx *gin.Context) {
	task := model.NewTask(t.db)

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "Task",
		"taskList": task.TaskList(),
	})
}

func (t *taskHandler) TaskStatusList(ctx *gin.Context) {
	list := []map[string]string{}
	for i, v := range model.TaskStatusLabels {
		list = append(list, map[string]string{
			"Id":    fmt.Sprint(i),
			"Label": v,
		})
	}
	ctx.JSON(http.StatusOK, list)
}

func (t *taskHandler) TaskListApi(ctx *gin.Context) {
	task := model.NewTask(t.db)
	ctx.JSON(http.StatusOK, task.TaskList())
}

func (t *taskHandler) TaskInput(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "input.tmpl", gin.H{
		"title": "Task Input",
	})
}

func (t *taskHandler) TaskCreateApi(ctx *gin.Context) {
	var form TaskForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "validate error")
		return
	}
	task := model.NewTask(t.db)
	sid, _ := ctx.Cookie("SESSION_ID")
	session := model.NewSession(t.redis)
	userId := session.GetUserId(sid)
	if err := task.TaskCreate(form.Name, userId, model.TASK_STATUS_TODO_ID); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "task: "+form.Name+" create")
}

func (t *taskHandler) TaskEditApi(ctx *gin.Context) {
	var form TaskFormEdit
	if err := ctx.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	task := model.NewTask(t.db)
	task.TaskUpdate(form.Id, form.Name, form.Status)
	ctx.JSON(http.StatusOK, nil)
}

func (t *taskHandler) TaskDeleteApi(ctx *gin.Context) {
	var form TaskFormDelete
	if err := ctx.ShouldBindUri(&form); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	task := model.NewTask(t.db)
	task.TaskDelete(form.Id)
	ctx.JSON(http.StatusOK, nil)
}
