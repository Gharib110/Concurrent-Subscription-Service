package handlers

import "net/http"

func (c *Config) LoadSessionManager(next http.Handler) http.Handler {
	return c.Session.LoadAndSave(next)
}
