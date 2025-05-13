package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// GoogleOAuthConfig is the OAuth2 config for Google
	GoogleOAuthConfig *oauth2.Config

	// Random string for OAuth2 state
	OAuthStateString = "random-state"
)

// UserInfo contains the information about the user that we get from Google
type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

// InitOAuth initializes the OAuth2 config
func InitOAuth() {
	// First try environment variables
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	// If environment variables are not set, attempt to read from credentials.json
	if clientID == "" || clientSecret == "" {
		// Read from credentials file
		data, err := os.ReadFile("credentials.json")
		if err == nil {
			var credentials struct {
				ClientID     string `json:"clientID"`
				ClientSecret string `json:"clientSecret"`
			}

			if err := json.Unmarshal(data, &credentials); err == nil {
				clientID = credentials.ClientID
				clientSecret = credentials.ClientSecret
			}
		}
	}

	// If we still don't have credentials, log a warning
	if clientID == "" || clientSecret == "" {
		fmt.Println("WARNING: Google OAuth credentials not found. Set GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET environment variables.")
		return
	}

	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// GoogleLoginHandler handles the Google login process
func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	if GoogleOAuthConfig == nil {
		http.Error(w, "OAuth not configured", http.StatusInternalServerError)
		return
	}

	url := GoogleOAuthConfig.AuthCodeURL(OAuthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleCallbackHandler processes the callback from Google
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if GoogleOAuthConfig == nil {
		http.Error(w, "OAuth not configured", http.StatusInternalServerError)
		return
	}

	state := r.FormValue("state")
	if state != OAuthStateString {
		fmt.Printf("Invalid oauth state, expected '%s', got '%s'\n", OAuthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := GoogleOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Get user info
	client := GoogleOAuthConfig.Client(r.Context(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Printf("Failed getting user info: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer response.Body.Close()

	var userInfo UserInfo
	if err = json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		fmt.Printf("Failed decoding response: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Save user info in session
	session, _ := GetSession(r)
	session.Values["authenticated"] = true
	session.Values["email"] = userInfo.Email
	session.Values["name"] = userInfo.Name
	session.Values["picture"] = userInfo.Picture
	session.Save(r, w)

	// Redirect to home page
	http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
}
