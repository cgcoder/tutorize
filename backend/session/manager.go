package session

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cgcoder/tutorize/backend/model"
	"github.com/cgcoder/tutorize/backend/utils"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	SESSION_KEY_PREFIX = "session:"
	SESSION_KEY_EXPIRE = 0 // 1 hour
)

type SessionManager interface {
	CreateSession(ctx context.Context, user model.SessionUser) (*model.Session, error)
	GetSession(ctx context.Context, sessionID string) (*model.Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
	UpdateSession(ctx context.Context, sessionID string, session *model.Session) error
}

type RedisSessionManager struct {
	client *redis.Client
}

func (r *RedisSessionManager) CreateSession(ctx context.Context, user model.SessionUser) (*model.Session, error) {
	session := &model.Session{
		Id:   uuid.New().String(),
		User: user,
	}
	jsonData, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	err = r.client.Set(ctx, SESSION_KEY_PREFIX+session.Id, jsonData, SESSION_KEY_EXPIRE).Err()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *RedisSessionManager) GetSession(ctx context.Context, sessionID string) (*model.Session, error) {
	session := &model.Session{}
	jsonData, err := r.client.Get(ctx, SESSION_KEY_PREFIX+sessionID).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonData), session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *RedisSessionManager) DeleteSession(ctx context.Context, sessionID string) error {
	err := r.client.Del(ctx, SESSION_KEY_PREFIX+sessionID).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisSessionManager) UpdateSession(ctx context.Context, sessionID string, session *model.Session) error {
	currSession, err := r.GetSession(ctx, sessionID)

	if currSession == nil || err != nil {
		return errors.New("session not valid")
	}

	jsonData, err := json.Marshal(session)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, SESSION_KEY_PREFIX+sessionID, jsonData, SESSION_KEY_EXPIRE).Err()
	if err != nil {
		return err
	}
	return nil
}

var sessionManager SessionManager

func InitSessionManager() error {
	redisHost, _ := utils.GetEnvOrDefault("REDIS_HOST", "")
	redisUser, _ := utils.GetEnvOrDefault("REDIS_USER", "")
	redisPassword, _ := utils.GetEnvOrDefault("REDIS_PASSWORD", "")
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Username: redisUser,
		Password: redisPassword,
		DB:       0, // use default DB
		Protocol: 2,
	})

	sessionManager = &RedisSessionManager{client: client}
	return nil
}

func GetSessionManager() SessionManager {
	return sessionManager
}
