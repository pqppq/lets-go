package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/pqppq/lets-go/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	// 1: month (January, Jan, 01, etc)
	// 2: day
	// 3: hour (15 is 3pm on a 24 hour clock)
	// 4: minute
	// 5: second
	// 6: year (2006)
	// 7: timezone (GMT-7 is MST)

	// DD Mon YYYY as HH:mm
	return t.Format("02 Jan 2006 at 15:04")
}
