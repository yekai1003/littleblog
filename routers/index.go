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

	note := dbs.Note{
		// Key:     "20101010",
		// UserID:  101,
		// Title:   "abc123",
		// Summary: "111",
		// Content: "1111",
		// Source:  "ancaaa",
		// Editor:  "1111",
		// Files:   "111",
		// Visit:   1,
		// Praise:  1010,
	}

	dataMap["totpage"] = 1
	dataMap["page"] = 1
	dataMap["title"] = "hehe"
	dataMap["Path"] = c.Request().RequestURI

	dataMap["notes"] = []dbs.Note{note}

	fmt.Println(os.Getwd())

	return c.Render(http.StatusOK, "index.html", dataMap)
}

//router: get /user
func GetUser(c echo.Context) error {
	dataMap := make(map[string]interface{})
	dataMap["title"] = "hehe"
	dataMap["Path"] = c.Request().RequestURI
	dataMap["IsLogin"] = true
	user := dbs.User{
		UserID: 1111,
		Name:   "yekai",
		Role:   0,
	}
	dataMap["User"] = user
	return c.Render(http.StatusOK, "user.html", dataMap)
}

//router: get /reg
func GetReg(c echo.Context) error {
	dataMap := make(map[string]interface{})
	// dataMap["title"] = "hehe"
	dataMap["Path"] = c.Request().RequestURI
	return c.Render(http.StatusOK, "reg.html", dataMap)
}

//router: get /message
func GetMessage(c echo.Context) error {
	dataMap := make(map[string]interface{})
	msg := []dbs.Message{}
	dataMap["messages"] = msg
	fmt.Println(c.Request().RequestURI)
	dataMap["Path"] = c.Request().RequestURI
	return c.Render(http.StatusOK, "message.html", dataMap)
}
