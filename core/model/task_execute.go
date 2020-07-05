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
	Context map[string]interface{} `gorm:"type:string;not null" json:"context"`
}

func (TaskExecute) TableName() string {
	return "task_execute"
}

func AddTaskExecute(TaskKey string, Params map[string]interface{}, Context map[string]interface{}) {
	core.Db.Create(&TaskExecute{
		Status:  core.TaskCreated,
		TaskKey: TaskKey,
		Params:  Params,
		Context: Context,
	})
}

// 更新任务执行记录状态
func (t *TaskExecute) UpdateTaskExecuteStatus(Status string) {
	taskExecute := GetTaskExecuteById(t.ID)
	core.Db.Model(&taskExecute).Update(TaskExecute{
		Status: Status,
	})
}

func GetTaskExecuteById(Id int64) (taskExecute TaskExecute) {
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
func PopTask() (TaskExecute, error) {
	// 队列中拿到任务，从表中查询该记录
	var taskExecute TaskExecute
	result := core.Redis.RPop(core.TaskQueueKey).Val()
	if result == "" {
		return taskExecute, nil
	}
	a := strings.Split(result, ":")
	Id, err := strconv.ParseInt(a[0], 10, 64)
	if err != nil {
		panic(err)
	}
	taskExecute = GetTaskExecuteById(Id)
	return taskExecute, err
}

// 向redis队列增加一个任务
func (t *TaskExecute) AddTaskToQueue() {
	task := fmt.Sprintf("%s:%s", strconv.FormatInt(t.BaseModel.ID, 10), t.TaskKey)
	core.Redis.LPush(core.TaskQueueKey, task)
}
