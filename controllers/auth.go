package controllers

import (
	"fmt"
	"net/http"

	"go-project/models"
	"go-project/utils/token"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]models.User)

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)
	str_user_id := fmt.Sprint(user_id)
	//fmt.Println("user_id_final", str_user_id)
	//fmt.Println("DB", db)
	data, _ := db[str_user_id]
	//b, _ := json.Marshal(data)
	//fmt.Println(string(b))
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
	//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	fmt.Println(err)
}

func Register(c *gin.Context) {

	var input models.AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	//u.BeforeSave()
	u.SaveUser(db)
	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func Login(c *gin.Context) {

	var input models.AuthInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	token, err := models.LoginCheck(u, db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
