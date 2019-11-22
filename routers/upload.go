package routers

import (
	"fmt"
	"io"

	"os"

	"github.com/labstack/echo"
)

type UploadFileInfo struct {
	Link string `json:"link"`
}

//@router post /upload/uploadfile
func UploadFile(c echo.Context) error {
	var resp ResponseData
	defer JsonWrite(c, &resp)
	resp.Code = RECODE_OK

	h, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Failed to FormFile", err)
		resp.Code = RECODE_PARAMERR
		return err
	}
	src, err := h.Open()
	defer src.Close()
	dstname := "upload/files/" + h.Filename
	dst, err := os.Create(dstname)
	if err != nil {
		fmt.Println("Failed to Create", err)
		resp.Code = RECODE_PARAMERR
		return err
	}

	resp.Data = UploadFileInfo{dstname}
	_, err = io.Copy(dst, src)
	return err
}

//@router post /upload/uploadfile
func UploadEditorFiles(c echo.Context) error {
	var resp ResponseData
	defer JsonWrite(c, &resp)
	resp.Code = RECODE_OK

	h, err := c.FormFile("files")
	if err != nil {
		fmt.Println("Failed to FormFile", err)
		resp.Code = RECODE_PARAMERR
		return err
	}
	src, err := h.Open()
	defer src.Close()
	dstname := "upload/files/" + h.Filename
	dst, err := os.Create(dstname)
	if err != nil {
		fmt.Println("Failed to Create", err)
		resp.Code = RECODE_PARAMERR
		return err
	}
	resp.Data = []string{dstname}
	_, err = io.Copy(dst, src)
	return err
}
