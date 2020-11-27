package main

import (
	"bytes"
	"fmt"
	"github.com/fdev-ci/golang-plugin-sdk/api"
	"github.com/fdev-ci/golang-plugin-sdk/log"
	"os/exec"
)
type bashParam struct {
	Script string `json:"script"`
}

func main() {
	log.Info("glang plugin bash starts")

	defer func() {
		if err := recover(); err != nil {
			log.Error("panic: ", err)
			api.FinishBuild(api.StatusError, "panic occurs")
		}
	}()

	run()

}

func run()  {
	// 打屏
	log.Info("\nBuildInfo:")
	log.Info("Project Name:     ", api.GetProjectDisplayName())
	log.Info("Pipeline Id:      ", api.GetPipelineId())
	log.Info("Pipeline Name:    ", api.GetPipelineName())
	log.Info("Pipeline Version: ", api.GetPipelineVersion())
	log.Info("Build Id:         ", api.GetPipelineBuildId())
	log.Info("Build Num:        ", api.GetPipelineBuildNumber())
	log.Info("Start Type:       ", api.GetPipelineStartType())
	log.Info("Start UserId:     ", api.GetPipelineStartUserId())
	log.Info("Start UserName:   ", api.GetPipelineStartUserName())
	log.Info("Start Time:       ", api.GetPipelineStartTimeMills())
	log.Info("Workspace:        ", api.GetWorkspace())

	// 输入参数解析到对象
	paramData := new(bashParam)
	api.LoadInputParam(paramData)
	log.Info(fmt.Sprintf("\nscript:%v", paramData.Script))

	// 业务逻辑
	log.Info("start build")
	err := build(paramData.Script)
	if err != nil {
		api.FinishBuildWithErrorCode(api.StatusError,err.Error(),1)
	} else {
		api.WriteOutput()
	}
	log.Info("build done")
}

func build(script string) error {

	cmd := exec.Command("sh", "-c",script)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Start()
	if err != nil {
		log.Info(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Info(err)
	}
	log.Info(out.String())
	return err

}