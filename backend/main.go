package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cgcoder/tutorize/backend/handlers"
	"github.com/cgcoder/tutorize/backend/model"
	"github.com/cgcoder/tutorize/backend/session"
	"github.com/cgcoder/tutorize/backend/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// validateSession handles session validation requests

// createSessionFromGoogle handles creating sessions from Google OAuth
func createSessionFromGoogle(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement actual Google OAuth verification
	// This would validate the OAuth token, extract user info, and create a session

	// For now, return a placeholder response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true, "session": {"id": "new-session-123", "user": {"id": "google-user", "name": "Google User"}}}`))
}

func main() {
	utils.ValidateEnv([]string{"SESSION_ENC_KEY", "REDIS_HOST", "REDIS_USER", "REDIS_PASSWORD"})
	session.InitSessionManager()
	sessionObj, err := session.GetSessionManager().CreateSession(context.TODO(), model.SessionUser{
		UserID:    "test-user",
		AccountId: "test-account",
		Name:      "Test User",
	})
	if err != nil {
		panic(err)
	}
	id1 := sessionObj.Id
	fmt.Println("Session created:", sessionObj)

	sessionObj, _ = session.GetSessionManager().CreateSession(context.TODO(), model.SessionUser{
		UserID:    "test-user",
		AccountId: "test-account",
		Name:      "Test User",
	})
	id2 := sessionObj.Id
	fmt.Println("Session created:", sessionObj, id2)
	sessionObj, err = session.GetSessionManager().GetSession(context.TODO(), id1)

	if err != nil {
		panic(err)
	}
	fmt.Println("Session retrieved:", sessionObj)

	// Create a new router
	r := chi.NewRouter()

	// Add middleware (optional)
	r.Use(middleware.Logger)          // Logs HTTP requests
	r.Use(middleware.Recoverer)       // Recovers from panics
	r.Use(handlers.SessionMiddleware) // Session middleware
	// Define the ping endpoint
	r.Get("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})

	// Define session-related routes
	//r.Get("/api/auth/session", getSession)
	//r.Post("/api/auth/session", createSession)
	handlers.RegisterSessionRoutes("/api/auth", r)
	// Start the server
	http.ListenAndServe(":5174", r)
}

func checkPrequisites() {
	model.CheckEncryptionKey()
}
