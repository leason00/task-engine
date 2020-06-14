package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-engine/api/e"
)

const StatusOK = http.StatusOK

func BaseResp(c *gin.Context, data interface{}, code int, msg string) {
	c.JSON(StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func MakeOkResp(c *gin.Context, data interface{}) {
	BaseResp(c, data, e.SUCCESS, "ok")
}

func MakeErrorResp(c *gin.Context, msg string, data interface{}) {
	BaseResp(c, data, e.ERROR, msg)
}

func MakePaginationResp(c *gin.Context, value interface{}, page int, perPage int, total int) {
	data := map[string]interface{}{
		"value":    value,
		"page":     page,
		"per_page": perPage,
		"total":    total,
	}
	c.JSON(StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}
