package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sihuayin/godist/models"
	"github.com/sihuayin/godist/pkg/e"
	"golang.org/x/crypto/bcrypt"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"user_name"`
	Password string `valid:"Required; MaxSize(50)" json:"user_password"`
}

type ChangePWD struct {
	OldPWD    string `json:"old_password"`
	NewPWD    string `json:"newpassword"`
	RepeatPWD string `json:"repeat_newpassword"`
}

// o := orm.NewOrm()
// err = o.Raw("SELECT * FROM `user` WHERE username= ?", userName).QueryRow(&user)
// beego.Info(user)
// err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
// if err != nil {
// 	c.SetJson(1, nil, "用户名或密码错误")
// 	return
// } else {
// 	if user.AuthKey == "" {
// 		userAuth := common.Md5String(user.Username + common.GetString(time.Now().Unix()))
// 		user.AuthKey = userAuth
// 		models.UpdateUserById(&user)
// 	}
// 	user.PasswordHash = ""
// 	c.SetJson(0, user, "")
// 	return
// }

func GetAuth(c *gin.Context) {

	var json auth

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("testing data,", json)
	valid := validation.Validation{}

	code := e.INVALID_PARAMS

	ok, _ := valid.Valid(&json)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": nil,
		})
		return
	}

	fmt.Println("查询", json.Username)
	// var user models.User
	user := models.FindOneByName(json.Username)

	if user != nil {
		fmt.Println(user.PasswordHash, json.Password)
		err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(json.Password))
		if err != nil {
			code = e.ERROR_AUTH
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": nil,
			})
			return
		}
	}

	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  e.GetMsg(code),
		"data": user,
	})
}

func AuthLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": "",
	})
}

func AuthChangePWD(c *gin.Context) {
	//哈希校验成功后 更新 auth_key
	var changePWD ChangePWD
	var err error
	if err = c.ShouldBindJSON(&changePWD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oldPassword := changePWD.OldPWD
	newPassword := changePWD.NewPWD
	repeatNewpassword := changePWD.RepeatPWD
	if oldPassword == "" || newPassword == "" || repeatNewpassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数有错误"})
		return
	}
	userInterface, ok := c.Get("User")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := userInterface.(*models.User)

	oldUser := models.FindOneByID(user.ID)
	if oldUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "查无此人"})
		return
	}
	//验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(oldUser.PasswordHash), []byte(oldPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "旧密码有误，请重新输入"})
		return
	} else {
		if newPassword == repeatNewpassword {
			password := []byte(newPassword)
			hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
			if err != nil {
				panic(err)
			}
			oldUser.PasswordHash = string(hashedPassword)
			models.UpdateUserById(oldUser)
			code := e.SUCCESS
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  e.GetMsg(code),
				"data": oldUser,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "两次密码输入不一致，请重新输入"})
			return
		}
	}
}
