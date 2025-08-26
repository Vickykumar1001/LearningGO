package handlers

import (
	"net/http"
	"session-23-gin-jwt/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {

	users, err := h.repo.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed To Get users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})

}
func (h *Handler) UpdateUsers(c *gin.Context) {

	idStr := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = idStr
	err := h.repo.UpdateUser(c, idStr, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Updated Succeesully",
		"user":    user,
	})

}

func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	err := h.repo.DeleteUser(c, idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
