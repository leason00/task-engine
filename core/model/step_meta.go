package model

// 任务流结构源数据
type StepMeta struct{
	Key string
	StepName string
	OrderID int
	TimeOut int  //单位s
}