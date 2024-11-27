package controllers

import (
	"net/http"
	"project/config"
	"project/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// ResetPassword allows users to update their password
func ResetPassword(c *gin.Context) {
	email := c.PostForm("email")
	newPassword := c.PostForm("newPassword")
	var user models.User

	// Find user and update password
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.HTML(http.StatusOK, "reset_password.html", gin.H{"error": "User does not exist!"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	config.DB.Save(&user)

	c.HTML(http.StatusOK, "reset_password.html", gin.H{"message": "Password updated successfully!"})
}
// RenderSignUp serves the SignUp HTML template
func RenderSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// RenderSignIn serves the SignIn HTML template
func RenderSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}
func ForgotPassword(c *gin.Context) {
	email := c.PostForm("email")
	favoriteColor := c.PostForm("favorite_color")
	newPassword := c.PostForm("new_password")
	var user models.User

	// Check if the user exists and their favorite color matches
	err := config.DB.Where("email = ? AND favorite_color = ?", email, favoriteColor).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			// Handle error differently for AJAX and regular requests
			if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or favorite color!"})
			} else {
				c.HTML(http.StatusBadRequest, "forgot_password.html", gin.H{
					"error": "Invalid email or favorite color!",
				})
			}
			return
		}
		// For unexpected errors
		if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred. Please try again."})
		} else {
			c.HTML(http.StatusInternalServerError, "forgot_password.html", gin.H{
				"error": "An unexpected error occurred. Please try again.",
			})
		}
		return
	}

	// Hash the new password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	config.DB.Save(&user)

	// Redirect or send success response
	if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
		c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully!"})
	} else {
		c.Redirect(http.StatusFound, "/signin")
	}
}


// SignUp handles user registration
func SignUp(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = string(hashedPassword)

	// Save the user in the database
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully! Please sign in."})
}

// SignIn handles user login
func SignIn(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "SignIn successful!"})
}
