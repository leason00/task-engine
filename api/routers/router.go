package routers

import (
	"task-engine/api/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) *gin.Engine {
	api := app.Group("/api/v1")
	{
		api.GET("/task_meta", v1.GetTaskMetas)

	}

	return app
}
