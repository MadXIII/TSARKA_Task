package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madxiii/tsarka_task/model"
)

func (h *API) GetUser(c *gin.Context) {
	param := c.Param("id")

	user, err := h.services.User.ByID(param)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (h *API) CreateUser(c *gin.Context) {
	var body model.User

	if err := c.BindJSON(&body); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.Create(body)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *API) UpdateUser(c *gin.Context) {
	var body model.User

	if err := c.BindJSON(&body); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	param := c.Param("id")

	err := h.services.User.Update(param, body)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "User updated",
	})
}

func (h *API) DeleteUser(c *gin.Context) {
	param := c.Param("id")

	err := h.services.User.Delete(param)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"staus": "User deleted",
	})
}
