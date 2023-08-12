package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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

func db() (*sql.DB, func()) {
	mysqlUser := "root"
	mysqlPassword := "root"
	mysqlHost := "127.0.0.1"
	dsn := fmt.Sprintf("%s:%s@(%s:3306)/task", mysqlUser, mysqlPassword, mysqlHost)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db, func() {
		db.Close()
	}
}
