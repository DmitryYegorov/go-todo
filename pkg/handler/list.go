package handler

import (
	todo "github.com/DmitryYegorov/go-todo/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "user not found")
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.CreateNewTodoList(input, userId.(int))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, ok := c.Get("userId")

	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "Lists not found")
		return
	}

	list, err := h.services.TodoList.GetAll(userId.(int))

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string][]todo.TodoList{
		"list": list,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, ok := c.Get("userId")
	listId, err := strconv.Atoi(c.Param("id"))

	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "Lists not found")
		return
	}

	list, err := h.services.TodoList.GetTodoListById(listId, userId.(int))

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]todo.TodoList{
		"list": list,
	})
}

func (h *Handler) updateList(c *gin.Context) {
	userId, ok := c.Get("userId")
	listId, err := strconv.Atoi(c.Param("id"))

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect listId")
		return
	}

	var input todo.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.UpdateList(listId, userId.(int), input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) deleteList(c *gin.Context) {
	userId, ok := c.Get("userId")
	listId, err := strconv.Atoi(c.Param("id"))

	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "incorrect listId")
		return
	}

	err = h.services.TodoList.DeleteListById(listId, userId.(int))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}
