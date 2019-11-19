package routers

import (
	"littleblog/cookies"
	"net/http"

	"github.com/labstack/echo"
)

// @router /new [get]
func NewPage(c echo.Context) error {
	dataMap := make(map[string]interface{})
	dataMap["key"] = 1001
	dataMap["Path"] = c.Request().RequestURI

	editor, _ := cookies.GetCookie(cookies.CookieName, "editor", c)
	if editor == "default" {
		editor = "note_new.html"
	} else {
		editor = "note_new2.html"
	}
	return c.Render(http.StatusOK, editor, dataMap)
}
