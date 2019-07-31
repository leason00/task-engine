package model

import (
	"github.com/leason00/task-engine/core"
)

// 任务队列信息
type TaskQueue struct{
	Id int
	Key string
}

// 从redis对列弹出一个任务
func (*TaskQueue) PopTask() (*TaskQueue, error) {

}

// 向redis对列增加一个任务
func (*TaskQueue) AddTask(id int, key string) (*TaskQueue, error) {
	//info := &TaskQueue{
	//	Id:          id,
	//	Key:      key,
	//}
	core.Redis.HSet()
}