package main

import (
	"fmt"
	"go/parser"
	"go/token"
	_ "task-engine/task-execute/tasks"
)

func main() {
	//for {
	//	time.Sleep(1 * time.Second)
	//	taskExecute, err := model.PopTask()
	//	if err != nil {
	//		core.Log.Error(err)
	//	}
	//	if taskExecute.TaskKey == "" {
	//		core.Log.Info("get nil task.")
	//		continue
	//	}
	//	fmt.Println(taskExecute.TaskKey)
	//	if taskExecute.Status != core.TaskWait {
	//		continue
	//	}
	//	base.Try(func() {
	//		taskExecute.UpdateTaskExecuteStatus(core.TaskDoing)
	//		core.Log.Info("begin execute task.")
	//	}, func(e interface{}) {
	//		core.Log.Error("doing execute task fail.")
	//		fmt.Println(e)
	//		taskExecute.UpdateTaskExecuteStatus(core.TaskFail)
	//	})
	//
	//}
	//t := reflect.TypeOf(tasks.TaskDemo{})
	//for i:=0;i<t.NumField();i++ {
	//	idx := t.Field(i)
	//	fmt.Println(idx.Name)
	//}
	//
	//pkg, err := importer.Default().Import("task-engine/task-execute/tasks")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	return
	//}
	//for _, declName := range pkg.Scope().Names() {
	//	fmt.Println(declName)
	//}
	//str, _ := os.Getwd()
	//fmt.Println(str)
	ParseFile("/Users/lixiangli/go/src/github.com/leason00/task-engine/task-execute/tasks/task_demo.go")
	//fmt.Printf("%#v", m)

}

func ParseFile(fileName string) (map[string]string, error) {
	fset := token.NewFileSet()
	// 解析文件，主要解析token
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	//fmt.Println(f)

	if err != nil {
		return nil, err
	}

	// 类型检查, 得到常量的值
	//conf := types.Config{Importer: importer.Default()}
	//pkg, err := conf.Check("/Users/lixiangli/go/src/github.com/leason00/task-engine/task-execute/tasks", fset, []*ast.File{f}, nil)
	//fmt.Println(pkg)
	fmt.Println(f.Scope.Objects)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for k, v := range f.Scope.Objects {
		fmt.Println(k)
		fmt.Println(v.Kind)

		//if v.Kind == ast.Con {
		//	d := v.Decl.(*ast.ValueSpec)
		//	m[pkg.Scope().Lookup(v.Name).(*types.Const).Val().String()] = d.Comment.Text()
		//}
	}
	return m, nil
}
