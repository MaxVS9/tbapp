package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi"
)

func (h *Handler) createList(c *gin.Context) {
	var input restapi.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.services.TodoList.Create(input)
}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
