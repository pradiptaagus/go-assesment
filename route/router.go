package route

import (
	"go-assesment/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	r.GET("/questions", controller.GetQuestions)
	r.POST("/questions", controller.StoreQuestion)
	r.GET("/questions/:id", controller.GetQuestion)
	r.PUT("/questions/:id", controller.UpdateQuestion)
	r.DELETE("/questions/:id", controller.DeleteQuestion)

	r.Run("localhost:8080")
}
