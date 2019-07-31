package model

import "github.com/satori/go.uuid"

type StepExecute struct{
	Id uuid.UUID
	TaskQueueId uuid.UUID
	StepMeta
	Status string
	Context map[string]interface{}
	StartTime string
	EndTime string
}