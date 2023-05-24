package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinitparekh17/go-mongo/db"
	"github.com/vinitparekh17/go-mongo/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	db.Init()
	id := c.Param("id")
	user, err := userModel.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	defer db.Database.Client().Disconnect(db.Ctx)
	c.Abort()
}

func (u UserController) Create(c *gin.Context) {
	db.Init()
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := userModel.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
	defer db.Database.Client().Disconnect(db.Ctx)
	c.Abort()
}
