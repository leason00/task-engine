package model

import (
	"task-engine/core"
)

type StepExecute struct {
	core.BaseModel

	TaskExecuteId int64  `gorm:"type:int64;not null" json:"task_execute_id"`
	TaskKey       string `gorm:"type:string;not null" json:"task_key"`
	StepName      string `gorm:"type:string;not null" json:"step_name"`
	OrderID       int    `gorm:"type:int;not null" json:"order_id"` // 顺序
	TimeOut       int    `gorm:"type:int;not null" json:"timeout"`  //单位s
	Status        string `gorm:"type:string;not null" json:"status"`
	StartTime     core.MyTime
	EndTime       core.MyTime
}

func (StepExecute) TableName() string {
	return "step_execute"
}

func AddStepExecute(TaskExecuteId int64, TaskKey string, StepName string) {
	// 根据task key 获取 task的meta信息
	stepMeta := GetbyTaskAndStep(TaskKey, StepName)
	core.Db.Create(&StepExecute{
		TaskExecuteId: TaskExecuteId,
		Status:        core.StepCreated,
		TaskKey:       TaskKey,
		StepName:      StepName,
		OrderID:       stepMeta.OrderID,
		TimeOut:       stepMeta.TimeOut,
	})
}

func GetById(Id int64) (stepExecute StepExecute) {
	core.Db.Where("id = ?", Id).First(&stepExecute)
	return
}

// 更新任务执行记录状态
func (t *StepExecute) UpdateStepExecuteStatus(Id int64, Status string) {
	stepExecute := GetById(Id)
	core.Db.Model(&stepExecute).Update(StepExecute{
		Status: Status,
	})
}

// 获取某个任务执行情况
func GetStepExecuteByTaskKey(TaskExecuteId int64, TaskKey string) (stepExecutes []StepExecute) {
	query := core.Db.Where("is_deleted = ?", 0)
	query.Where("task_execute_id = ? AND task_key = ?", TaskExecuteId, TaskKey).Order("order_id desc").Find(&stepExecutes)
	return
}
