package cookies

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const CookieName = "blogcookie"

func SetCookie(cookiename string, keyval map[string]string, c echo.Context) error {
	for k, v := range keyval {
		cookie := new(http.Cookie)
		cookie.Name = cookiename + "-" + k
		cookie.Value = v
		c.SetCookie(cookie)
	}
	return nil
}

func GetCookie(cookiename string, key string, c echo.Context) (string, error) {
	cookie, err := c.Cookie(cookiename + "-" + key)
	if err != nil {
		fmt.Println("GetCookie err", err)
		return "", err
	}

	return cookie.Value, nil
}
