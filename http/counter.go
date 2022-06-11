package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *API) CountGet(c *gin.Context) {
	res, err := h.services.Count.CounterGet()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count is": res,
	})
}

func (h *API) CountAdd(c *gin.Context) {
	param := c.Param("number")

	err := h.services.Count.CounterAdd(param)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "added",
	})
}

func (h *API) CountSub(c *gin.Context) {
	param := c.Param("number")

	err := h.services.Count.CounterSub(param)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "substracted",
	})
}
