package controller

import (
	"fmt"
	"go-assesment/helper"
	"go-assesment/message"
	"go-assesment/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	var questions []model.Question
	helper.DB().Find(&questions)
	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: message.StatusOK,
		Data:    questions,
	})
}

func GetQuestion(c *gin.Context) {
	var question model.Question
	var id = c.Param("id")

	result := helper.DB().Find(&question, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: message.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	fmt.Println(result.RowsAffected)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  http.StatusNotFound,
			Message: message.StatusNotFound,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: message.StatusOK,
		Data:    question,
	})
}

func StoreQuestion(c *gin.Context) {
	var question model.Question

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: message.StatusBadRequest,
			Data:    err.Error(),
		})
		return
	}

	result := helper.DB().Create(&question)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: message.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	// helper.DB().Where("question = ?", question.Question).Find(&question)

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: message.StatusOK,
		Data:    question,
	})
}

func UpdateQuestion(c *gin.Context) {
	var question model.Question
	var id = c.Param("id")

	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: message.StatusBadRequest,
			Data:    err.Error(),
		})
		return
	}

	var _question model.Question
	helper.DB().Find(&_question, id)

	if _question.ID == 0 {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  http.StatusNotFound,
			Message: message.StatusNotFound,
			Data:    nil,
		})
		return
	}

	result := helper.DB().Model(&model.Question{}).Where("id", id).Update("question", question.Question)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: message.StatusInternalServerError,
			Data:    nil,
		})
		return
	}

	helper.DB().Find(&question, id)

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: message.StatusOK,
		Data:    question,
	})
}

func DeleteQuestion(c *gin.Context) {
	var question model.Question
	var id = c.Param("id")

	result := helper.DB().Find(&question, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: message.StatusInternalServerError,
			Data:    nil,
		})
		return
	} else if question.ID == 0 {
		c.JSON(http.StatusNotFound, model.Response{
			Status:  http.StatusNotFound,
			Message: message.StatusNotFound,
			Data:    nil,
		})
		return
	}

	_result := helper.DB().Delete(&model.Question{}, id)

	if _result.Error != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  http.StatusInternalServerError,
			Message: message.StatusInternalServerError,
			Data:    nil,
		})
	}

	c.JSON(http.StatusOK, model.Response{
		Status:  http.StatusOK,
		Message: message.StatusOK,
		Data:    nil,
	})
}
