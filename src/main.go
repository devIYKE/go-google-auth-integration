package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"googleauth/auth"
	"googleauth/handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// Store for session management
	sessionStore *sessions.CookieStore
)

func main() {
	// Initialize the session store
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		// Use a default key for development
		sessionKey = "google-signin-example-session-key"
		fmt.Println("WARNING: Using default session key. Set SESSION_KEY environment variable in production.")
	}
	sessionStore = sessions.NewCookieStore([]byte(sessionKey))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}

	// Initialize the auth package with our session store
	auth.InitSessionStore(sessionStore)

	// Initialize auth service
	auth.InitOAuth()
	// Create router
	r := mux.NewRouter()
	// Static files
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	r.HandleFunc("/profile", handlers.ProfileHandler).Methods("GET")

	// Auth routes
	r.HandleFunc("/auth/google/login", auth.GoogleLoginHandler).Methods("GET")
	r.HandleFunc("/auth/google/callback", auth.GoogleCallbackHandler).Methods("GET")
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Ikechukwu Samuel Madu's Google Authentication Application")
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
