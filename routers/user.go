//router: get /user
package routers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const (
	RECODE_OK         = "0"
	RECODE_DBERR      = "4001"
	RECODE_SESSIONERR = "4002"
	RECODE_LOGINERR   = "4003"
	RECODE_PARAMERR   = "4004"
	RECODE_UNKNOWERR  = "4101"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库操作错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_UNKNOWERR:  "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}

	return recodeText[RECODE_UNKNOWERR]
}

type ResponseData struct {
	Code    string      `json:"code"`
	CodeMsg string      `json:"codemsg"`
	Action  string      `json:"action"`
	Count   string      `json:"count"`
	Data    interface{} `json:"data"`
}

func JsonWrite(c echo.Context, r *ResponseData) error {
	r.CodeMsg = RecodeText(r.Code)
	return c.JSON(http.StatusOK, r)
}

type RegRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

//router: post /reg
func Reg(c echo.Context) error {
	var resp ResponseData
	var userreq RegRequest
	defer JsonWrite(c, &resp)
	resp.Code = RECODE_OK
	err := c.Bind(&userreq)
	if err != nil {
		resp.Code = RECODE_PARAMERR
		fmt.Println("Failed to echo.Bind", err)
		return err
	}
	fmt.Println(userreq)
	if userreq.Email == "" ||
		userreq.Name == "" ||
		userreq.Password == "" {
		fmt.Println("Params err")
		resp.Code = RECODE_PARAMERR
		return err
	}

	if userreq.Password != userreq.Password2 {
		fmt.Println("passwd must same")
		return err
	}
	resp.Action = "/user"
	return nil
}

//router: post /login
func Login(c echo.Context) error {
	var resp ResponseData
	var userreq RegRequest
	defer JsonWrite(c, &resp)
	resp.Code = RECODE_OK
	err := c.Bind(&userreq)
	if err != nil {
		resp.Code = RECODE_PARAMERR
		fmt.Println("Failed to echo.Bind", err)
		return err
	}
	if userreq.Email == "" || userreq.Password == "" {
		resp.Code = RECODE_LOGINERR
		err = errors.New("user or password err")
		fmt.Println("", err)
		return err
	}
	resp.Action = "/"
	return nil
}

// @router /setting/editor [post]

func Editor(c echo.Context) error {
	var resp ResponseData
	resp.Code = RECODE_OK
	defer JsonWrite(c, &resp)
	return nil
}
