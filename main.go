package main

import (
	"fmt"
	"html/template"
	"io"
	"littleblog/routers"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	fmt.Println("name===", name, data)
	// w.Write([]byte("hello world"))
	// return nil
	err := t.templates.ExecuteTemplate(w, name, data)
	fmt.Println("err === ", err)
	return err
}

var funcMap = template.FuncMap{
	"add":      add,
	"date":     date,
	"str2html": str2html,
}

func add(x, y int) int {
	return x + y
}

func date(t time.Time, format string) string {
	return ""
}
func str2html(content string) string {
	return ""
}

func main() {
	e := echo.New()

	e.Static("/static", "static")
	e.GET("/", routers.Index)

	t := &TemplateRenderer{
		templates: template.Must(template.New("").Funcs(funcMap).ParseFiles("views/index.html", "views/user.html", "views/reg.html", "views/setting.html", "views/details.html")),
	}
	// 赋值
	e.Renderer = t

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:8080"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
	e.Use(middleware.Logger())
	e.Start(":8080")
}
