package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type session struct {
	redis *redis.Client
}

func NewSession(redis *redis.Client) *session {
	return &session{redis}
}

type userSession struct {
	UserId int64 `json:"user_id"`
}

func (s *session) Create(userId int64, ctx *gin.Context) {
	sid := uuid.New().String()
	userJson, err := json.Marshal(userSession{userId})
	if err != nil {
		panic("json marshal error")
	}
	ctx.SetCookie("SESSION_ID", sid, 3600*3, "/", "localhost", true, true)
	if err = s.redis.Set(context.Background(), sid, string(userJson), time.Hour*3).Err(); err != nil {
		panic("redis set error")
	}
}

func (s *session) GetUserId(sid string) int64 {
	d, _ := s.redis.Get(context.Background(), sid).Result()
	var us userSession
	json.Unmarshal([]byte(d), &us)
	return us.UserId
}
