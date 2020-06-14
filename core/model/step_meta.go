package model

import (
	"task-engine/core"
)

// 任务流结构源数据
type StepMeta struct {
	core.BaseModel

	TaskKey  string `gorm:"type:string;not null" json:"task_key"`
	StepName string `gorm:"type:string;not null" json:"step_name"`
	OrderID  int    `gorm:"type:int;not null" json:"order_id"` // 顺序
	TimeOut  int    `gorm:"type:int;not null" json:"timeout"`  //单位s
}

func (StepMeta) TableName() string {
	return "step_meta"
}

func GetbyTaskAndStep(TaskKey string, StepName string) (stepMeta StepMeta) {
	core.Db.Where("TaskKey = ? AND StepName= ?", TaskKey, StepName).First(&stepMeta)
	return
}
