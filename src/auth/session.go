package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// Key for session store
	sessionStore *sessions.CookieStore
)

// InitSessionStore initializes the session store
func InitSessionStore(store *sessions.CookieStore) {
	sessionStore = store
}

// GetSession returns the current session
func GetSession(r *http.Request) (*sessions.Session, error) {
	return sessionStore.Get(r, "google-auth-session")
}

// IsAuthenticated checks if the user is authenticated
func IsAuthenticated(r *http.Request) bool {
	session, err := GetSession(r)
	if err != nil {
		return false
	}

	auth, ok := session.Values["authenticated"].(bool)
	return ok && auth
}

// RequireLogin is a middleware that redirects to login if not authenticated
func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next(w, r)
	}
}
