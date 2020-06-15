package main

import (
	"fmt"
	"task-engine/core"
	"time"
)
import "task-engine/core/model"

// 定时从mysql表中读取待执行的任务
// 放入待消费的队列
// 队列的生产者

func main() {
	for {
		time.Sleep(1 * time.Second)
		var result bool
		result = core.Db.HasTable(&model.TaskExecute{})
		fmt.Println(result)

		var maps map[string]interface{}
		maps = make(map[string]interface{})
		maps["Status"] = core.TaskCreated
		taskExecutes := model.GetTaskExecute(1, 10, maps)
		for _, taskExecute := range taskExecutes {
			fmt.Println(taskExecute.TaskKey)
			taskExecute.AddTaskToQueue()
			taskExecute.UpdateTaskExecuteStatus(core.TaskWait)
		}
	}
}
