package helpers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
    Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
   return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo) {
    tmpl := template.Must(template.ParseGlob("views/*.html"))
    t := newTemplate(tmpl)
    e.Renderer = t
}

func newTemplate(templates *template.Template) echo.Renderer {
    return &Template{
        Templates: templates,
    }
}
