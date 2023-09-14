package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"pbi-btpns-api/internal/exception"
	"pbi-btpns-api/internal/model"
	"pbi-btpns-api/internal/utils"
)

type UploadController interface {
	UploadPhoto(c *gin.Context)
}

type uploadControllerImpl struct{}

func (u *uploadControllerImpl) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		panic(exception.InvariantError{Msg: "image not found"})
	}

	extension := utils.GetFileExtension(file.Filename)

	// check file extension is permitted
	isPermitted := isFilePermitted(extension)
	if !isPermitted {
		panic(exception.InvariantError{Msg: "image extension not permitted, only : png, jpg, or jpeg"})
	}

	id := uuid.New().String()
	filePath := fmt.Sprintf("./public/uploads/photos/%s.%s", id, extension)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		panic(err)
	}

	c.JSON(200, model.WebResponse{
		Status:  model.Success,
		Code:    200,
		Message: "Yay, success to add new photo",
		Data: model.UploadPhotoResponse{
			Url: fmt.Sprintf("/static/photos/%s.%s", id, extension),
		},
	})
}

func isFilePermitted(extension string) bool {
	suffix := []string{"png", "jpg", "jpeg"}
	for _, v := range suffix {
		if extension == v {
			return true
		}
	}

	return false
}
