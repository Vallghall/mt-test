package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrSamePerson      = errors.New("person A and person B are the same person")
	ErrBadRequest      = errors.New("request body has invalid parameters")
	ErrInvalidCommType = errors.New("invalid communication type")
)

func catch(c *gin.Context, err error, status int) {
	if err != nil {
		c.AbortWithStatusJSON(status, map[string]string{
			"message": err.Error(),
		})
	}
}
