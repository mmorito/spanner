package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mmorito/spanner/apis/apis/helloworld"
)

// type Template struct {
// 	templates *template.Template
// }

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

// Router is router of API
func Router(e *echo.Echo) {

	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("templates/*.html")),
	// }
	// e.Renderer = t
	// e.GET("/", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "index.html", nil)
	// })

	g := e.Group("/api")

	// hello
	ping := g.Group("/hello")
	ping.GET("", helloworld.Get)

}
