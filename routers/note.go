package routers

import (
	"bytes"
	"fmt"
	"littleblog/cookies"
	"littleblog/dbs"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)

// @router /new [get]
func NewPage(c echo.Context) error {
	dataMap := make(map[string]interface{})
	dataMap["key"] = dbs.Dbx.GetNoteKey()
	dataMap["Path"] = c.Request().RequestURI

	editor, _ := cookies.GetCookie(cookies.CookieName, "editor", c)
	if editor == "default" {
		editor = "note_new.html"
	} else {
		editor = "note_new2.html"
	}
	//保存到数据库
	user_id, _ := cookies.GetCookie(cookies.CookieName, "userid", c)
	userid, _ := strconv.ParseInt(user_id, 10, 64)
	dbs.Dbx.SaveNoteKey(userid, dataMap["key"].(string))

	return c.Render(http.StatusOK, editor, dataMap)
}

/*
	files := ctx.GetString("files", "")
	summary, _ := getSummary(content)
	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
*/

func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))

	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	strRune := []rune(str)
	if len(strRune) > 400 {
		strRune = strRune[:400]
	}
	return string(strRune) + "...", nil
}

//@router /save/:key [post]
func NoteSave(c echo.Context) error {
	var resp ResponseData
	resp.Code = RECODE_OK
	defer JsonWrite(c, &resp)
	n := dbs.Note{}
	n.Key = c.Param("key")
	//user_id, _ := cookies.GetCookie(cookies.CookieName, "userid", c)

	fmt.Println("key ===", n.Key)

	err := c.Bind(&n)
	if err != nil {
		resp.Code = RECODE_PARAMERR
		fmt.Println("Failed to echo.Bind", err)
		return err
	}
	// n.Title = noteMap["title"]
	// n.Editor = "default"
	// n.Content = noteMap["content"]
	n.Summary, _ = getSummary(n.Content)
	fmt.Println(n)
	if err = dbs.Dbx.SaveNoteInfo(n); err != nil {
		resp.Code = RECODE_DBERR
		return err
	}
	resp.Action = "/details/" + n.Key

	return nil
}
