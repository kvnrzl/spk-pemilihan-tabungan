package app

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// func VerifyToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token := c.Request.Header.Get("Authorization")
// 		if token == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"code":  http.StatusUnauthorized,
// 				"error": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}
// 		token = strings.Replace(token, "Bearer ", "", 1)
// 		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return []byte("secret"), nil
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"code":  http.StatusUnauthorized,
// 				"error": "Unauthorized",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
