package api

import (
	"fmt"
	"github.com/alexandrevicenzi/unchained"
	"github.com/gin-gonic/gin"
	"github.com/kayes-shawon/go-gin/pkg/db"
	"github.com/kayes-shawon/go-gin/pkg/models"
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

}

func UserLoginRefresh(c *gin.Context) {

}
