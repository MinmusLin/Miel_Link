package router

import (
	con "backend/controller"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/register", con.Register)
	r.POST("/login", con.Login)
	r.POST("/logout", con.Logout)
	r.POST("/getInfo", middleware.JWTAuthMiddleware(), con.GetInfo)
	r.POST("/uplink", middleware.JWTAuthMiddleware(), con.Uplink)
	r.POST("/getProductInfo", con.GetProductInfo)
	r.POST("/getProductList", middleware.JWTAuthMiddleware(), con.GetProductList)
	r.POST("/getAllProductInfo", middleware.JWTAuthMiddleware(), con.GetAllProductInfo)
	r.POST("/getProductHistory", middleware.JWTAuthMiddleware(), con.GetProductHistory)
	r.POST("/ipfsUpload", con.IpfsUpload)
	r.POST("/ipfsDownload", con.IpfsDownload)
	r.GET("/download", con.Download)
	return r
}
