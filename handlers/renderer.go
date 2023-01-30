package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var templatePath = "./templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
}

func (c *Config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", templatePath),
		fmt.Sprintf("%s/header.partial.gohtml", templatePath),
		fmt.Sprintf("%s/navbar.partial.gohtml", templatePath),
		fmt.Sprintf("%s/footer.partial.gohtml", templatePath),
		fmt.Sprintf("%s/alerts.partial.gohtml", templatePath),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", templatePath, t))

	for _, v := range partials {
		templateSlice = append(templateSlice, v)
	}

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		c.ErrLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, c.AddDefaultData(td, r))
	if err != nil {
		c.ErrLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = c.Session.PopString(r.Context(), "flash")
	td.Warning = c.Session.PopString(r.Context(), "warning")
	td.Error = c.Session.PopString(r.Context(), "error")

	if c.IsAuthenticated(r) {
		td.Authenticated = true
	}

	td.Now = time.Now()

	return td
}

func (c *Config) IsAuthenticated(r *http.Request) bool {
	return c.Session.Exists(r.Context(), "userID")
}
