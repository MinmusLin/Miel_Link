package controller

import (
	"backend/model"
	"backend/pkg"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.MysqlUser
	user.UserID = pkg.GenerateID()
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.RealInfo = pkg.EncryptByMD5(c.PostForm("username"))
	err := pkg.InsertUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Register failed: " + err.Error(),
		})
		return
	}
	var args []string
	args = append(args, user.UserID)
	args = append(args, c.PostForm("userType"))
	args = append(args, user.RealInfo)
	res, err := pkg.ChaincodeInvoke("RegisterUser", args)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Register failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Register successfully",
		"txid":    res,
	})
}

func Login(c *gin.Context) {
	var user model.MysqlUser
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	var err error
	user.UserID, err = pkg.GetUserID(user.Username)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "未找到该用户",
		})
		return
	}
	userType, err := GetUserType(user.UserID)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Login failed: " + err.Error(),
		})
		return
	}
	err = pkg.Login(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Login failed: " + err.Error(),
		})
		return
	}
	jwt, err := pkg.GenToken(user.UserID, userType)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Login failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Login successfully",
		"jwt":     jwt,
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Logout successfully",
	})
}

func GetUserType(userID string) (string, error) {
	userType, err := pkg.ChaincodeQuery("GetUserType", userID)
	if err != nil {
		return "", err
	}
	return userType, nil
}

func GetInfo(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(200, gin.H{
			"message": "Get user type failed",
		})
	}
	userType, err := GetUserType(userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Get user type failed: " + err.Error(),
		})
	}
	username, err := pkg.GetUsername(userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Get user type failed: " + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":     200,
		"message":  "Get user type successfully",
		"userType": userType,
		"username": username,
	})
}
