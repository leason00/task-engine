package main

import (
	"fmt"
	"task-engine/core"
	"task-engine/core/model"
	"task-engine/task-execute/base"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		taskExecute, err := model.PopTask()
		if err != nil {
			core.Log.Error(err)
		}
		if taskExecute.TaskKey == "" {
			core.Log.Info("get nil task.")
			continue
		}
		fmt.Println(taskExecute.TaskKey)
		if taskExecute.Status != core.TaskWait {
			continue
		}
		base.Try(func() {
			taskExecute.UpdateTaskExecuteStatus(core.TaskDoing)
			core.Log.Info("begin execute task.")
		}, func(e interface{}) {
			core.Log.Error("doing execute task fail.")
			fmt.Println(e)
			taskExecute.UpdateTaskExecuteStatus(core.TaskFail)
		})

	}

}
