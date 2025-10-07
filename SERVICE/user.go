package SERVICE

import (
	"Golang_task4/CONF"
	"Golang_task4/DAO"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册
func UserRegister(c *gin.Context) error {
	var user DAO.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return err
	}
	user.Password = string(hashedPassword)

	if err := DAO.Conn.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存用户失败"})
		return err
	}
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
	return nil
}

// 登录
func UserLogin(c *gin.Context) error {
	// 登录用户
	var user DAO.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	// 数据库的用户
	var dbUser DAO.User
	if result := DAO.Conn.Where("username = ?", user.Username).First(&dbUser); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "查询用户失败"})
		return result.Error
	}

	// 验证密码
	compareErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if compareErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return compareErr
	}
	// 生成JWT token
	token, err := CONF.GenerateJWT(dbUser.Username, dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": dbUser.ID,
	})
	return nil
}
