package task_dispatch

import "github.com/leason00/task-engine/core"

// 定时从mysql表中读取待执行的任务
// 放入待消费的对列
// 对列的生产者

func TaskDispatch(){
	rs, err := core.DB.Query("select id, task_key from task_queue where status='wait_exec'")
	if err != nil {
		core.Log.Error(err)
	}
	for rs.Next() {
		var id int
		var task_key string

		err = rs.Scan(&id, &task_key)
		if err != nil {
			core.Log.Error(err)
		}
		Redis.HSet()
	}
}