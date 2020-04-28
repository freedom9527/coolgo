package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type AdminController struct {
	BaseController
	Db *gorm.DB
}

func (a *AdminController)IndexEndpoint(c *gin.Context)  {
	a.Prepare(c)
	a.pageData =  make(map[string]interface{})
	a.pageData["message"] = a.controller+"OK1!"
	c.JSON(http.StatusOK, a.pageData)
}

func (a *AdminController)DeleteEndpoint(c *gin.Context)  {
	contrl := c.Param("controller")
	c.JSON(http.StatusOK, gin.H{
		"message": contrl+"Delete!",
	})
}

func (a *AdminController)EditEndpoint(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "Edit!",
	})
}