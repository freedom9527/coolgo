package controllers

import (
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	controller string             //controller name
	action     string             //action name
	pageData   map[string]interface{}
}

func (b *BaseController)Prepare(c *gin.Context)  {
	b.controller = c.Param("controller")
}

