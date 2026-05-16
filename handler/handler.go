package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"wiza.core/domain"
)

type ClientHandler struct {
	svc domain.ClientService
}

func NewClientHandler(svc domain.ClientService) *ClientHandler {
	return &ClientHandler{svc: svc}
}

func (h *ClientHandler) GetByIIN(c *gin.Context) {
	iin := c.Param("iin")

	client, err := h.svc.GetByIIN(c.Request.Context(), iin)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, client)
}
