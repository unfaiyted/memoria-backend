// constants/config.go
package constants

// DefaultConfig represents the default configuration values
var DefaultConfig = map[string]interface{}{
	// App defaults
	"app.name":        "Memoria",
	"app.environment": "development",
	"app.appURL":      "http://localhost:3000",
	"app.apiBaseURL":  "http://localhost:8080",
	"app.logLevel":    "info",
	"app.maxPageSize": 100,

	// Database defaults
	"db.host":     "localhost",
	"db.port":     "5432",
	"db.name":     "memoria",
	"db.user":     "postgres",
	"db.password": "yourpassword",
	"db.maxConns": 20,
	"db.timeout":  30,

	// HTTP defaults
	"http.port":             "8080",
	"http.readTimeout":      30,
	"http.writeTimeout":     30,
	"http.idleTimeout":      60,
	"http.enableSSL":        false,
	"http.rateLimitEnabled": true,
	"http.requestsPerMin":   100,

	// Auth defaults
	"auth.enableLocal":     true,
	"auth.sessionTimeout":  60,
	"auth.enable2FA":       false,
	"auth.tokenExpiration": 24,
	"auth.allowedOrigins":  []string{"http://localhost:5173"},
}
