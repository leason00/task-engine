package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task-engine/api/utils"
	"task-engine/core/model"
)

func GetTaskMetas(c *gin.Context) {
	data := make(map[string]interface{})

	page := c.DefaultQuery("page", "1")
	perPage := c.DefaultQuery("per_page", "10")
	taskKey := c.DefaultQuery("task_key", "")
	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	var maps map[string]interface{}
	maps = make(map[string]interface{})
	if taskKey != "" {
		maps["task_key"] = taskKey
	}
	taskMetas, total := model.GetTaskMeta(pageInt, perPageInt, maps)
	data["value"] = taskMetas
	data["total"] = total

	utils.MakeOkResp(c, data)
}
