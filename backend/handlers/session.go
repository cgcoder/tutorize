package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cgcoder/tutorize/backend/model"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func RegisterSessionRoutes(prefix string, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {
		r.Get("/session", getSession)
		r.Post("/session/create_from_creds", createSessionFromCreds)
		r.Post("/session/google", createSessionFromGoogle)
	})
}

func SessionMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic can be added here
		sessionIdCookie, _ := r.Cookie(model.SESSION_COOKIE)
		var cookieId string
		var session *model.Session = &model.Session{}
		if sessionIdCookie != nil {
			cookieId = sessionIdCookie.Value
		} else {
			cookieId = uuid.NewString()
		}
		session.Id = cookieId
		r = r.WithContext(context.WithValue(r.Context(), model.SESSION, session))
		http.SetCookie(w, &http.Cookie{
			Name:     model.SESSION_COOKIE,
			Value:    cookieId,
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour)})
		handler.ServeHTTP(w, r)
	})
}

func createSessionFromCreds(w http.ResponseWriter, r *http.Request) {

}

func createSessionFromGoogle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func getSession(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(model.SESSION).(*model.Session)
	if session == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
