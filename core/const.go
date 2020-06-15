package core

const TaskQueueKey = "task_engine_queue"

const task_queue = "wait_exec"

const (
	TaskCreated = "created"
	TaskWait    = "wait"
	TaskDone    = "done"
	TaskDoing   = "doing"
	TaskFail    = "fail"
)

const (
	StepCreated = "created"
	StepWait    = "wait"
	StepDone    = "done"
	StepDoing   = "doing"
)
