package v1

import (
	"eshop/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) isUser(c *gin.Context) {

	userType, err := getUserType(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if userType != domain.UserTypeUser {
		newErrorResponse(c, http.StatusUnauthorized, "You are not a user")
		return
	}

}
