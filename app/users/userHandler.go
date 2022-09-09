package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerRegisterUser(c *gin.Context) {
	var user UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userC := NewUserController()
	userSaved, err := userC.CreateUser(user)
	if err != nil {
		log.Println("got failed when register user ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": userSaved})
}
