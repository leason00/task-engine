package model

import (
	"fmt"
	"github.com/leason00/task-engine/core"
	"strconv"
	"strings"
)

// 任务队列
type TaskQueueBase struct{
	Id int
	Key string
}

// 任务队列详细信息
type TaskQueue struct{
	TaskQueue
}

func NewTaskQueue(id int, key string) *TaskQueueBase {
	return &TaskQueueBase{
		Id: id,
		Key: key,
	}
}

// 从redis队列弹出一个任务
func (t *TaskQueueBase) PopTask() (*TaskQueueBase, error) {
	result := core.Redis.RPop(core.TaskQueueKey).Val()
	a := strings.Split(result, ":")
	id, err := strconv.Atoi(a[0])
	if err != nil {
		return nil, err
	}
	return NewTaskQueue(id, a[0]), nil
}

// 向redis对列增加一个任务
func (t *TaskQueueBase) AddTask(id int, key string) {
	task := fmt.Sprintf("%s:%s", strconv.Itoa(id), key)
	core.Redis.LPush(core.TaskQueueKey, task)
}


// 从mysql获取待执行的任务
func (t *TaskQueue) GetTasks(count int)  (*TaskQueueBase, error) {

}