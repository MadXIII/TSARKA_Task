package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madxiii/tsarka_task/model"
)

func (h *API) EmailSearch(c *gin.Context) {
	var body model.Check

	if err := c.BindJSON(&body); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Search.CheckEmail(body.Email)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}
