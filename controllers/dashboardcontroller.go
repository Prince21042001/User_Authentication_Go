package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RenderDashboard renders the dashboard page
func RenderDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"technologies": []string{"Go", "Gin", "PostgreSQL", "HTML", "CSS", "JavaScript"},
		"libraries":    []string{"bcrypt", "gorm", "fetch API"},
		"database":     "PostgreSQL",
	})
}
