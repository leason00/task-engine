package model

import (
	"task-engine/core"
)

// 任务流结构源数据
type TaskMeta struct {
	core.BaseModel

	TaskKey string `gorm:"type:string;not null" json:"task_key"`
}

func (TaskMeta) TableName() string {
	return "task_meta"
}

// 获取任务流
func GetTaskMeta(pageNum int, pageSize int, maps map[string]interface{}) ([]TaskMeta, int) {
	var taskMeta []TaskMeta
	var total int
	query := core.Db.Where("is_deleted = ?", 0).Where(maps)
	query.Count(&total)
	query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("created_at desc").Find(&taskMeta)
	return taskMeta, total
}
