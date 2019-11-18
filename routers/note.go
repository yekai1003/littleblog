package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

// @router /new [get]
func NewPage(c echo.Context) error {
	dataMap := make(map[string]interface{})
	dataMap["key"] = 1001
	dataMap["Path"] = c.Request().RequestURI
	return c.Render(http.StatusOK, "note_new2.html", dataMap)
}
