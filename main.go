package main

import (
	"fmt"
	"html/template"
	"io"
	"littleblog/routers"
	"os"
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

	err := t.templates.ExecuteTemplate(w, name, data)

	fmt.Println("err === ", err, ",name ==", name)
	return err
}

var funcMap = template.FuncMap{
	"add":      add,
	"date":     date,
	"str2html": str2html,
	"equrl":    equrl,
}

func add(x, y int) int {
	return x + y
}

func date(t time.Time, format string) string {
	return t.Format(format)
}
func str2html(content string) string {
	return ""
}

func equrl(s1, s2 string) bool {
	return s1 == s2
}

func getAllFiles(dirname string) ([]string, error) {
	var files []string
	fd, err := os.Open(dirname)
	if err != nil {
		fmt.Println("Failed to Open", err)
		return files, err
	}
	infos, err := fd.Readdir(-1)
	for _, k := range infos {
		if k.IsDir() {
			continue
		}
		files = append(files, dirname+"/"+k.Name())
	}
	//fmt.Println(files)
	return files, nil
}

func main() {
	e := echo.New()

	e.Static("/static", "static")
	e.GET("/", routers.Index)
	e.GET("/user", routers.GetUser)
	e.GET("/reg", routers.GetReg)
	e.GET("/message", routers.GetMessage)
	e.POST("/reg", routers.Reg)
	e.POST("/login", routers.Login)
	e.POST("/setting/editor", routers.Editor)
	e.GET("/note/new", routers.NewPage)
	e.POST("/upload/uploadfile/", routers.UploadFile)
	files, _ := getAllFiles("views")
	t := &TemplateRenderer{
		templates: template.Must(template.New("").Funcs(funcMap).ParseFiles(files...)),
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
