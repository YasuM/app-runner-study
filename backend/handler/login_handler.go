package handler

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"task-app-study/entity"
	"task-app-study/model"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type loginHandler struct {
	db    *sql.DB
	redis *redis.Client
}

type userSession struct {
	UserId int64 `json:"user_id"`
}

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginHandler(db *sql.DB, redis *redis.Client) *loginHandler {
	return &loginHandler{db, redis}
}

func (l *loginHandler) Login(ctx *gin.Context) {
	var form LoginForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, "validate error")
		return
	}
	query := entity.New(l.db)

	user, err := query.GetUserdByEmail(context.Background(), form.Email)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "login error")
		return
	}
	ctx.SetSameSite(http.SameSiteNoneMode)
	session := model.NewSession(l.redis)
	session.Create(user.ID, ctx)
	ctx.JSON(http.StatusOK, "login success")
}
