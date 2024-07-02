package controller

import (
	"context"
	"log"
	"net/http"
	"toy/service"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	workspace := c.Query("workspace")
	email := c.Query("email")

	googleInfo, err := service.GetTenantUserInfoAndGoogleInfo(context.Background(), email, workspace)

	if err != nil {
		log.Printf("Error retrieving Google info: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving Google info"})
		return
	}

	if googleInfo != nil {
		c.JSON(http.StatusOK, googleInfo)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Google info found for the given user ID"})
	}
}
