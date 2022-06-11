package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madxiii/tsarka_task/model"
)

func (h *API) Calculate(c *gin.Context) {
	var body model.Hash

	if err := c.BindJSON(&body); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hashID, err := h.services.Hash.CalculateBody(body.Input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": hashID,
	})
}

func (h *API) GetResult(c *gin.Context) {
	param := c.Param("id")

	res, err := h.services.Hash.GetResult(param)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if res < 1 {
		c.JSON(http.StatusCreated, gin.H{
			"result": "PENDING",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}
