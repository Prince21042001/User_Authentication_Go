package routes

import (
	"net/http"
	"project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// HTML templates
	router.LoadHTMLGlob("templates/*")

	// Routes
	router.GET("/signin", controllers.RenderSignIn)
	router.GET("/signup", controllers.RenderSignUp)
	router.GET("/forgot-password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forgot_password.html", nil)
	})
	router.GET("/dashboard", controllers.RenderDashboard)
	router.GET("/reset-password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reset_password.html", nil)
	})

	router.POST("/signin", controllers.SignIn)
	router.POST("/signup", controllers.SignUp)
	router.POST("/forgot-password", controllers.ForgotPassword)
	router.POST("/reset-password", controllers.ResetPassword)
}
