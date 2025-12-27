package settings

import "net/http"

type Application struct {
	Config AppConfig
}

type AppConfig struct {
	Server ServerConfig
	// Database DatabaseConfig
	// Redis RedisConfig
}

func GetConfig() AppConfig {
	config := AppConfig{
		Server: loadServerConfig(),
		// Database: loadDatabaseConfig(),
		// Redis: loadRedisConfig(),
	}
	return config
}

func (a *Application) Start() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.Config.Server.Handler,
	}

	server.ListenAndServe()
}
