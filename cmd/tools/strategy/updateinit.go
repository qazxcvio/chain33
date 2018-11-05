package strategy

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gitlab.33.cn/chain33/chain33/cmd/tools/tasks"
)

type updateInitStrategy struct {
	strategyBasic

	consRootPath  string
	dappRootPath  string
	storeRootPath string
}

func (this *updateInitStrategy) Run() error {
	fmt.Println("Begin run chain33 update init.go.")
	defer fmt.Println("Run chain33 update init.go finish.")
	if err := this.initMember(); err != nil {
		return err
	}
	return this.runImpl()
}

func (this *updateInitStrategy) initMember() error {
	var outputPath string
	gopath := os.Getenv("GOPATH")
	if len(gopath) > 0 {
		outputPath = filepath.Join(gopath, "/src/gitlab.33.cn/chain33/chain33/plugin/")
	}
	if len(outputPath) <= 0 {
		return errors.New("Chain33 Plugin Not Existed")
	}
	this.consRootPath = fmt.Sprintf("%s/consensus/", outputPath)
	this.dappRootPath = fmt.Sprintf("%s/dapp/", outputPath)
	this.storeRootPath = fmt.Sprintf("%s/store/", outputPath)
	return nil
}

func (this *updateInitStrategy) runImpl() error {
	var err error
	task := this.buildTask()
	for {
		if task == nil {
			break
		}
		err = task.Execute()
		if err != nil {
			mlog.Error("Execute command failed.", "error", err, "taskname", task.GetName())
			break
		}
		task = task.Next()
	}
	return err
}

func (this *updateInitStrategy) buildTask() tasks.Task {
	taskSlice := make([]tasks.Task, 0)
	taskSlice = append(taskSlice,
		&tasks.UpdateInitFileTask{
			Folder: this.consRootPath,
		},
		&tasks.UpdateInitFileTask{
			Folder: this.dappRootPath,
		},
		&tasks.UpdateInitFileTask{
			Folder: this.storeRootPath,
		},
	)

	task := taskSlice[0]
	sliceLen := len(taskSlice)
	for n := 1; n < sliceLen; n++ {
		task.SetNext(taskSlice[n])
		task = taskSlice[n]
	}
	return taskSlice[0]
}
