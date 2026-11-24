package main

import "os"

// appConfig holds settings supplied through the environment.
// TODO: Add Firebase project settings here when the database is configured.
type appConfig struct {
	Address string
}

func loadConfig() appConfig {
	address := os.Getenv("APP_ADDRESS")
	if address == "" {
		address = ":8080"
	}
	return appConfig{Address: address}
}
