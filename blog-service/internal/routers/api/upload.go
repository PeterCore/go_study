package api

import (
	"my-blog-sevice/global"
	"my-blog-sevice/internal/service"
	"my-blog-sevice/pkg/app"
	"my-blog-sevice/pkg/convert"
	"my-blog-sevice/pkg/errcode"
	"my-blog-sevice/pkg/upload"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UpLoadFiles(c *gin.Context) {
	response := app.NewResponse(c)
	form, err := c.MultipartForm()
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	files := form.File["file"]
	svc := service.New(c.Request.Context())
	urls := make([]string, 0)
	size := c.Request.ContentLength
	for _, fileHeader := range files {
		if fileHeader == nil || fileType < 0 {
			response.ToErrorResponse(errcode.InvalidParams)
			return
		}
		fileInfo, err := svc.UpLoadFile(upload.FileType(fileType), size, fileHeader)
		if err != nil {
			global.Logger.Errorf("svc.UploadFile err : %v", err)
			errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		urls = append(urls, fileInfo.AccessUrl)

	}
	response.ToResponse(gin.H{
		"file_access_urls": urls,
	})

}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	_, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	if fileHeader == nil || fileType < 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	size := c.Request.ContentLength
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UpLoadFile(upload.FileType(fileType), size, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err : %v", err)
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
