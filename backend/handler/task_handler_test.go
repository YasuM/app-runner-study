package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task-app-study/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
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

	h := NewTaskHandler(db, newRedisClient())
	h.TaskCreate(ctx)

	fmt.Println(w.Code)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "\"task: "+name+" create\"")
}

func TestTaskListApi(t *testing.T) {
	db, dbClose := db()
	defer dbClose()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("GET", "/", nil)
	now := time.Now()
	_, err := db.Exec("insert into task(name, status, created_at) values(?, ?, ?)", "task1", model.TASK_STATUS_TODO_ID, now.Format("2006-01-02 15:04:05"))
	if err != nil {
		t.Fatal(err)
	}
	now = now.Add(time.Hour)
	_, err = db.Exec("insert into task(name, status, created_at) values(?, ?, ?)", "task2", model.TASK_STATUS_TODO_ID, now.Format("2006-01-02 15:04:05"))
	if err != nil {
		t.Fatal(err)
	}

	h := NewTaskHandler(db, newRedisClient())
	h.TaskList(ctx)

	var actual []model.TaskEntity
	json.Unmarshal(w.Body.Bytes(), &actual)
	assert.Equal(t, "task2", actual[0].Name)
	assert.Equal(t, model.TaskStatusLabels[actual[0].Status], actual[0].StatusLabel)
	assert.Equal(t, "task1", actual[1].Name)
	assert.Equal(t, model.TaskStatusLabels[actual[1].Status], actual[1].StatusLabel)
}

func init() {
	txdb.Register("txdb", "mysql", "root:root@/task_test?parseTime=true")
}

func db() (*sql.DB, func()) {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		panic(err)
	}
	return db, func() {
		db.Close()
	}
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
