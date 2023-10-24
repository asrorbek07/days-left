package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDays(c *gin.Context) {
	daysLeft, err := h.service.CalculateDaysCount()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//return days left in json format
	c.JSON(http.StatusOK, map[string]interface{}{
		"days_left": daysLeft,
	})
}
