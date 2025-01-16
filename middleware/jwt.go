package middleware

import (
	"demo/libs"
	"github.com/gin-gonic/gin"
	"time"
)

var Jwt *libs.JwtStruct

func init() {
	Jwt = libs.NewJwt()
}

func JwtGuard() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "未携带",
			})
			c.Abort()
			return
		}

		claims, err := Jwt.VerifyToken(authHeader)

		if err != nil {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "解析失败",
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.JSON(200, gin.H{
				"code":    400,
				"message": "过期了",
			})
			c.Abort()
			return
		}

		c.Set("id", claims.Id)
		c.Next()
	}
}
