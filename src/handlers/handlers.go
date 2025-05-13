package handlers

import (
	"net/http"
	"text/template"

	"googleauth/auth"
)

// HomeHandler renders the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html", "templates/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authenticated := auth.IsAuthenticated(r)
	data := map[string]interface{}{
		"Title":         "Home - Ikechukwu Samuel Madu's Google Auth",
		"Authenticated": authenticated,
	}

	if authenticated {
		session, _ := auth.GetSession(r)
		data["Name"] = session.Values["name"]
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// LoginHandler renders the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("templates/login.html", "templates/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"Title": "Login - Ikechukwu Samuel Madu's Google Auth",
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// LogoutHandler handles user logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.GetSession(r)
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ProfileHandler renders the user profile page
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	session, _ := auth.GetSession(r)
	data := map[string]interface{}{
		"Title":         "Profile - Ikechukwu Samuel Madu's Google Auth",
		"Authenticated": true,
		"Name":          session.Values["name"],
		"Email":         session.Values["email"],
		"Picture":       session.Values["picture"],
	}
	tmpl, err := template.ParseFiles("templates/profile.html", "templates/layout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}
