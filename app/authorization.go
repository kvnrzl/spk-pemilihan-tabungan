package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		// token = strings.Replace(token, "Bearer ", "", 1)
		// _, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// 	}
		// 	return []byte("secret"), nil
		// })

		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"code":  http.StatusUnauthorized,
		// 		"error": "Unauthorized",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}
