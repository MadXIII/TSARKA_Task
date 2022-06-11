package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madxiii/tsarka_task/model"
)

func (h *API) Find(c *gin.Context) {
	var body model.Find

	if err := c.BindJSON(&body); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res := h.services.Finder.CheckStr(body.Substr)

	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}
