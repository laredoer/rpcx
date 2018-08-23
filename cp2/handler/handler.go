package handler

import (
	"thresher/cp2/model"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	u := model.User
	c.JSON(200, gin.H{
		"message": u.Get(1, 2),
	})
}
