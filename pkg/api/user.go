package api

import (
	"fmt"
	"github.com/alexandrevicenzi/unchained"
	"github.com/gin-gonic/gin"
	"github.com/kayes-shawon/go-gin/pkg/db"
	"github.com/kayes-shawon/go-gin/pkg/models"
	"github.com/kayes-shawon/go-gin/pkg/utils"
	"net/http"
)

func CreateUser(c *gin.Context) {
	user := &models.User{}
	//raw, _ := c.GetRawData()
	_ = c.BindJSON(&user)
	dbCon := db.ConnectDB()

	hash, err := unchained.MakePassword(user.Password, "", "default")
	if err != nil {
		return
	}
	user.Password = hash

	_, err = dbCon.Model(user).Insert()
	if err != nil {
		fmt.Printf("Error\n", err)
		return
	}


	data := map[string]interface{} {
		"id" : user.Id,
		"username": user.UserName,
	}

	c.JSON(200, data)
}

func UserLogin(c *gin.Context) {
	user := &models.User{}
	_ = c.BindJSON(&user)
	dbCon := db.ConnectDB()
	password := user.Password
	err := dbCon.Model(user).Where("user_name = ?", user.UserName).Select()
	if err != nil {
		return
	}
	encodedPassword := user.Password
	valid, err := unchained.CheckPassword(password, encodedPassword)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
	}

	payload := map[string]interface{} {
		"username": user.UserName,
	}
	token, err := utils.Encode(payload)

}

func UserLoginRefresh(c *gin.Context) {

}
