package handlers

import "net/http"

func (c *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "home.page.gohtml", nil)
}
