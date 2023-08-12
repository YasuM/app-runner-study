package handler

import (
	"app-runner-study/model"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestTaskCreateApi(t *testing.T) {
	db, dbClose := db()
	defer dbClose()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	name := "task1"
	form := TaskForm{Name: name}
	j, _ := json.Marshal(form)
	ctx.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(j))

	h := NewTaskHandler(db)
	h.TaskCreateApi(ctx)

	fmt.Println(w.Code)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "\"task: "+name+" create\"")
}

type A struct {
	Name        string
	Status      int
	StatusLabel string
}

func TestTaskListApi(t *testing.T) {
	db, dbClose := db()
	defer dbClose()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("GET", "/", nil)
	now := time.Now()
	db.Exec("insert into task(name, status, created_at) values(?, ?, ?)", "task1", model.TASK_STATUS_TODO_ID, now.Format("2006-01-02T15:04:05Z07:00"))
	now = now.Add(time.Hour)
	db.Exec("insert into task(name, status, created_at) values(?, ?, ?)", "task2", model.TASK_STATUS_TODO_ID, now.Format("2006-01-02T15:04:05Z07:00"))

	h := NewTaskHandler(db)
	h.TaskListApi(ctx)

	var actual []model.TaskEntity
	json.Unmarshal(w.Body.Bytes(), &actual)
	assert.Equal(t, "task2", actual[0].Name)
	assert.Equal(t, model.TaskStatusLabels[actual[0].Status], actual[0].StatusLabel)
	assert.Equal(t, "task1", actual[1].Name)
	assert.Equal(t, model.TaskStatusLabels[actual[1].Status], actual[1].StatusLabel)
}

func db() (*sql.DB, func()) {
	txdb.Register("txdb", "mysql", "root:root@/task_test")
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		panic(err)
	}
	return db, func() {
		db.Close()
	}
}
