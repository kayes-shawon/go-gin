package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	//token, _ := c.Get("Authorization")
	//tk := token.(string)
	//
	//if len(tk) < 1 {
	//	c.String(200, "No token found")
	//}
	fmt.Println("Middleware")
	c.Next()

}
