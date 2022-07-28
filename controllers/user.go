package controllers

import (
	"crud-golang/middlewares"
	"crud-golang/models"
	"crud-golang/repo"
	"crud-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	// _ "github.com/go-sql-driver/mysql"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func AuthHandler(c *gin.Context) {
	// var login repo.Login
	var user models.User
	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&user); binderr != nil {
		log.Error().Err(binderr).
			Msg("Error occurred while binding request data")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}
	vu, err := repo.Login(&user)
	if err != nil {
		
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse params" + err.Error(),
			"data":   nil,
		})
	}

	if vu == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "User not found",
			"data":   nil,
		})
	}

	if vu != nil {
		tokenString, _ := middlewares.CreateToken(vu.Username)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "Success",
			"data": gin.H{"token": tokenString},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Verified Failed.",
	})
	return
}

func CreateUser(c *gin.Context) {
	var user models.User
	// Bind request payload with our model
	if binderr := c.ShouldBindJSON(&user); binderr != nil {
		log.Error().Err(binderr).
			Msg("Error occurred while binding request data")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": binderr.Error(),
		})
		return
	}
	user.FillDefaults()

	hashedPassword, hashErr := Hash(user.Password)
	if hashErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": hashErr,
		})
		return
	}

	user.Password = string(hashedPassword)

	userInfo, err := repo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userInfo,
		"user": user,
	})
}

func GetAllUsers(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)
	var user models.User
	userLists, totalRows, err := repo.GetAllUsers(&user, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data":      userLists,
		"totalList": totalRows,
	})

}

// func GetUserById(c *gin.Context) {

// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.FindUserById(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func UpdateUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.FindUserById(&user, id)
// 	if err != nil {
// 		println("no data found")
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	c.BindJSON(&user)
// 	err = Model.UpdateUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func DeleteUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.DeleteUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// }
