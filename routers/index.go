package routers

import (
	"fmt"
	"littleblog/cookies"
	"littleblog/dbs"
	"net/http"
	"os"
	"strconv"

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
	dataMap["IsLogin"] = false

	name, err := cookies.GetCookie(cookies.CookieName, "name", c)

	if err != nil {
		fmt.Println("user is 游客")
		//return err
		user := dbs.User{
			Name: "游客",
			Role: 1,
		}
		dataMap["User"] = user
	} else {
		user := dbs.User{
			Name: name,
		}
		user.Email, _ = cookies.GetCookie(cookies.CookieName, "email", c)
		user.Editor, _ = cookies.GetCookie(cookies.CookieName, "editor", c)
		role, _ := cookies.GetCookie(cookies.CookieName, "role", c)
		user.Role, _ = strconv.Atoi(role)
		//user.Role = int(role64)

		dataMap["User"] = user
		dataMap["IsLogin"] = true
	}

	return c.Render(http.StatusOK, "user.html", dataMap)
}

//router: get /reg
func GetReg(c echo.Context) error {
	dataMap := make(map[string]interface{})
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

// @router: get /setting
func GetSetting(c echo.Context) error {
	return c.Render(http.StatusOK, "setting.html", "")
}

func GetDetail(c echo.Context) error {
	dataMap := make(map[string]interface{})
	key := c.Param("key")
	note := dbs.Note{Key: key}
	err := dbs.Dbx.GetNoteByKey(&note)
	if err != nil {
		return err
	}

	user, err := dbs.Dbx.GetUserByID(note.UserID)
	if err != nil {
		fmt.Println("Failed to GetUserByID", err, note.UserID)
		return err
	}

	dataMap["praise"] = false

	// messages, _ := c.Dao.QueryMessageForNote(note.Key)
	// c.Data["messages"] = messages
	dataMap["note"] = note
	dataMap["user"] = user
	dataMap["Path"] = c.Request().RequestURI
	return c.Render(http.StatusOK, "details.html", dataMap)
}
