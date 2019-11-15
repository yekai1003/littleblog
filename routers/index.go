package routers

import (
	"fmt"
	"littleblog/dbs"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

//router: get /
func Index(c echo.Context) error {

	dataMap := make(map[string]interface{})

	ns := dbs.Note{
		Key:     "20101010",
		UserID:  101,
		Title:   "abc123",
		Summary: "111",
		Content: "1111",
		Source:  "ancaaa",
		Editor:  "1111",
		Files:   "111",
		Visit:   1,
		Praise:  1010,
	}

	dataMap["totpage"] = 1
	dataMap["page"] = 1
	dataMap["title"] = "hehe"

	dataMap["notes"] = ns

	fmt.Println(os.Getwd())

	return c.Render(http.StatusOK, "index.html", dataMap)
}
