package app

import (
	"github.com/gin-gonic/gin"
	"pbi-btpns-api/internal/exception"
	"pbi-btpns-api/internal/model"
)

func ErrorHandler(c *gin.Context, err any) {
	if v, ok := err.(exception.ValidationError); ok {
		c.JSON(400, model.WebResponse{
			Status:  model.Fail,
			Code:    400,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.JsonParseError); ok {
		c.JSON(400, model.WebResponse{
			Status:  model.Fail,
			Code:    400,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.InvariantError); ok {
		c.JSON(400, model.WebResponse{
			Status:  model.Fail,
			Code:    400,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.AuthenticationError); ok {
		c.JSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	if v, ok := err.(exception.AuthorizationError); ok {
		c.JSON(401, model.WebResponse{
			Status:  model.Fail,
			Code:    401,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}
	if v, ok := err.(exception.NotFoundError); ok {
		c.JSON(404, model.WebResponse{
			Status:  model.Fail,
			Code:    404,
			Message: v.Msg,
			Data:    nil,
		})
		return
	}

	// INTERNAL SERVER ERROR
	// TODO write error log file

	c.JSON(500, model.WebResponse{
		Status:  model.Error,
		Code:    500,
		Message: "Internal Server Error",
		Data:    nil,
	})
}
