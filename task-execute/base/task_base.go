package base

import (
	"task-engine/core"
	"task-engine/core/model"
)

type TaskBase struct {
	model.StepExecute
}

func (b *TaskBase) TaskExecute(StepName string) {
	Try(func() {
		core.Log.Info("begin execute task.")
	}, func(e interface{}) {
		core.Log.Error("doing execute task fail.")
		panic(e)
	})
}
