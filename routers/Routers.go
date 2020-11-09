package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Helper for api
	uc := hUser()

	// API route for version version-1
	v1 := r.Group("/api/v1")

	v1.POST("signup", uc.SignUp)

	return r
}