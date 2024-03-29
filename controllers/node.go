package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jianggushi/topstory/models"
)

// ListNodes .
func ListNodes(c *gin.Context) {
	nodes, err := models.GetNodes()
	if err != nil {
		c.JSON(404, formatJSON(nil, err))
		return
	}
	c.JSON(200, formatJSON(nodes, nil))
}

// GetNodeByID .
func GetNodeByID(c *gin.Context) {
	var arg struct {
		ID int `uri:"id" binding:"required"`
	}
	err := c.ShouldBindUri(&arg)
	if err != nil {
		c.JSON(400, formatJSON(nil, err))
		return
	}
	node, err := models.GetNodeByID(arg.ID)
	if err != nil {
		c.JSON(404, formatJSON(nil, err))
		return
	}
	c.JSON(200, formatJSON(node, nil))
}
