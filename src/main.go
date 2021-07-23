package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Mobile string `json:"mobile" binding:"required,numeric,len=11"`
	Name   string `json:"name" binding:"required,min=5"`
	Email  string `json:"email" binding:"required,email,min=5"`
}

func (u User) String() string {
	return fmt.Sprintf("[Mobile=%s,Name=%s,Email=%s]", u.Mobile, u.Name, u.Email)
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		c.JSON(200, user)
	})

	r.Run(":8080")
}
