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

func (app *Application) Start() {

	serverConfig:=app.Config.Server

	server := &http.Server{
		Addr:    ":8080",
		Handler: serverConfig.Handler,
		ReadTimeout: serverConfig.ReadTimeout,
		WriteTimeout: serverConfig.WriteTimeout,
		IdleTimeout: serverConfig.IdleTimeout,
	}

	server.ListenAndServe()
}
