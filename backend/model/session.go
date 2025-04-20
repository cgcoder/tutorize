package model

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type SessionUser struct {
	UserID    string `json:"user_id"`
	AccountId string `json:"account_id"`
	Name      string `json:"name"`
}

type Session struct {
	Id    string      `json:"id"`
	User  SessionUser `json:"session"`
	Valid bool        `json:"valid"`
}
type ContextType int

const (
	SESSION_COOKIE string = "cid"
)

const (
	SESSION ContextType = iota
)

// getEncryptionKey returns the key used for AES256 encryption
// Retrieved from SESSION_ENC_KEY environment variable
func getEncryptionKey() []byte {
	key := os.Getenv("SESSION_ENC_KEY")
	if key == "" {
		// Fallback for development only
		// In production, you should ensure the env var is set and handle this case properly
		fmt.Println("WARNING: SESSION_ENC_KEY environment variable not set, PANIC!")
		os.Exit(2)
	}
	return []byte(key)
}

func GetSessionFromJwt(tokenStr string) Session {
	// This function should implement JWT verification logic
	// For now, we will just return a dummy response

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return getEncryptionKey(), nil
	})

	if err != nil || !token.Valid {
		return Session{}
	}

	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return Session{}
	}
	sessionId := tokenClaims["session_id"].(string)
	userId := tokenClaims["user_id"].(string)
	accountId := tokenClaims["account_id"].(string)
	first_name := tokenClaims["first_name"].(string)

	return Session{
		Id:   sessionId,
		User: SessionUser{UserID: userId, AccountId: accountId, Name: first_name},
	}
}
func CheckEncryptionKey() {
	// panic!
	getEncryptionKey()
}
