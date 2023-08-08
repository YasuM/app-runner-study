package handler

import (
	"app-runner-study/model"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	db *sql.DB
}

type TaskForm struct {
	Name string `form:"task" json:"task" binding:"required"`
}

func NewTaskHandler(db *sql.DB) *taskHandler {
	return &taskHandler{db}
}

func (t *taskHandler) TaskList(ctx *gin.Context) {
	task := model.NewTask(t.db)

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "Task",
		"taskList": task.TaskList(),
	})
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

func (t *taskHandler) TaskCreate(ctx *gin.Context) {
	var form TaskForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.HTML(http.StatusOK, "input.tmpl", gin.H{
			"title": "Task Input",
			"error": err.Error(),
		})
		return
	}
	task := model.NewTask(t.db)
	task.TaskCreate(form.Name)

	ctx.Redirect(http.StatusMovedPermanently, "/task")
}

func (t *taskHandler) TaskCreateApi(ctx *gin.Context) {
	var form TaskForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	task := model.NewTask(t.db)
	task.TaskCreate(form.Name)
	ctx.JSON(http.StatusOK, nil)
}
