# Ikechukwu Samuel Madu's Go Web Application with Google Sign-In

This is a Go web application developed by Ikechukwu Samuel Madu that demonstrates how to integrate Google Sign-In authentication using OAuth 2.0.

## Project Structure

```
.
├── src/
│   ├── auth/         # Authentication related code
│   │   ├── oauth.go   # OAuth2 configuration and handlers
│   │   └── session.go # Session management
│   ├── handlers/     # HTTP request handlers
│   ├── models/       # Data models
│   ├── static/       # Static assets (CSS, JS)
│   ├── templates/    # HTML templates
│   ├── main.go       # Application entry point
│   └── config.go     # Configuration
└── README.md         # Project documentation
```

## Prerequisites

- Go 1.16 or higher
- A Google Cloud Platform account with OAuth 2.0 credentials

## Setup

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/go-google-auth.git
   cd go-google-auth
   ```

2. Configure your Google OAuth credentials:
   - Create a project in the [Google Cloud Console](https://console.cloud.google.com/)
   - Navigate to "APIs & Services" > "Credentials"
   - Configure the OAuth consent screen with necessary information
   - Create OAuth 2.0 credentials (Web application type)
   - Add `http://localhost:8080/auth/google/callback` as an authorized redirect URI

3. Configure the application:
   - Copy `credentials.json.template` to `credentials.json`
   - Edit `credentials.json` and add your Google OAuth credentials:
     ```json
     {
       "clientID": "YOUR_CLIENT_ID",
       "clientSecret": "YOUR_CLIENT_SECRET"
     }
     ```
   
   Alternatively, you can set environment variables:
   - `GOOGLE_CLIENT_ID`: Your Google Client ID
   - `GOOGLE_CLIENT_SECRET`: Your Google Client Secret
   - `SESSION_KEY`: Secret key for session encryption (optional, a default will be used if not provided)

4. Install dependencies:
   ```
   cd src
   go mod tidy
   ```

5. Run the application:
   ```
   go run .
   ```

6. Open your browser and navigate to [http://localhost:8080](http://localhost:8080)

## Features

- Google Sign-In authentication
- User session management
- Protected routes requiring authentication
- User profile information display
- Responsive design

## Technologies Used

- Go programming language
- Gorilla Mux (HTTP router)
- Gorilla Sessions (session management)
- Google OAuth2 API for authentication
- HTML/CSS/JavaScript for the frontend

## Development

To run the application in development mode:

```
cd src
go run .
```

The application will be available at [http://localhost:8080](http://localhost:8080).

## License

MIT
