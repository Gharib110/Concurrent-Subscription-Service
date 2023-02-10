package handlers

import "net/http"

func (c *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "home.page.gohtml", nil)
}

func (c *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "home.page.gohtml", nil)
}

func (c *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = c.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		c.ErrLog.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := c.Data.User.GetByEmail(email)
	if err != nil {
		c.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		c.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if !validPassword {
		c.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	c.Session.Put(r.Context(), "userID", user.ID)
	c.Session.Put(r.Context(), "user", user)

	c.Session.Put(r.Context(), "flash", "Successful Login")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *Config) Logout(w http.ResponseWriter, r *http.Request) {
	_ = c.Session.Destroy(r.Context())
	_ = c.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (c *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	c.render(w, r, "register.page.gohtml", nil)
}

func (c *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (c *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {

}
