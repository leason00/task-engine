package model

import (
	"fmt"
	"strconv"
	"strings"
	"task-engine/core"
)

// 任务队列
type TaskExecute struct {
	core.BaseModel
	Status  string `gorm:"type:string;not null" json:"status"`
	TaskKey string `gorm:"type:string;not null" json:"task_key"`
	Params  map[string]interface{}
}

func (TaskExecute) TableName() string {
	return "task_execute"
}

func AddTaskExecute(TaskKey string, Params map[string]interface{}) {
	core.Db.Create(&TaskExecute{
		Status:  core.TaskCreated,
		TaskKey: TaskKey,
		Params:  Params,
	})
}

// 更新任务执行记录状态
func (t *TaskExecute) UpdateTaskExecuteStatus(Id int64, Status string) {
	taskExecute := t.GetById(Id)
	core.Db.Model(&taskExecute).Update(TaskExecute{
		Status: Status,
	})
}

func (TaskExecute) GetById(Id int64) (taskExecute TaskExecute) {
	core.Db.Where("id = ?", Id).First(&taskExecute)
	return
}

func GetTaskExecute(pageNum int, pageSize int, maps map[string]interface{}) (taskExecutes []TaskExecute) {
	fmt.Println(maps)
	query := core.Db.Where("is_deleted = ?", 0)
	query.Where(maps).Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&taskExecutes)
	return
}

// 从redis队列弹出一个任务
func (t *TaskExecute) PopTask() (TaskExecute, error) {
	// 队列中拿到任务，从表中查询该记录
	result := core.Redis.RPop(core.TaskQueueKey).Val()
	a := strings.Split(result, ":")
	Id, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		panic(err)
	}
	return t.GetById(Id), err
}

// 向redis队列增加一个任务
func (t *TaskExecute) AddTaskToQueue() {
	task := fmt.Sprintf("%s:%s", strconv.FormatInt(t.BaseModel.ID, 10), t.TaskKey)
	core.Redis.LPush(core.TaskQueueKey, task)
}
