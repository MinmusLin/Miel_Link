package middleware

import (
    "backend/pkg"

    "github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
    return func(c *gin.Context) {
        authHeader := c.Request.Header.Get("Authorization")
        if authHeader == "" {
            c.JSON(200, gin.H{
                "code": 401,
                "msg":  "JWT Token authentication failed",
            })
            c.Abort()
            return
        }
        mc, err := pkg.ParseToken(authHeader)
        if err != nil {
            c.JSON(200, gin.H{
                "code": 401,
                "msg":  "JWT Token authentication failed",
                "data": err.Error(),
            })
            c.Abort()
            return
        }
        c.Set("userID", mc.UserID)
        c.Next()
    }
}
