/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-21 20:48:46
 * @LastEditTime: 2022-02-07 20:54:22
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/gosible/executer.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */

package gosible

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/0xe8998e/gosible/ssh"

	"github.com/goinggo/mapstructure"
	"github.com/pkg/errors"
)

const (
	FAILED    = "failed"
	SUCESSFUL = "sucessful"
)

type Executer struct {
	SshConfig *ssh.Config
}
type ExecuterParam struct {
	Host Host
	Task Task
}

type ExecuterResult struct {
	HostResult HostResult
	Error      error
}

type ShellParam struct {
	Data string
}

type CopyParam struct {
	Src string
	Dst string
}

type TemplateParam struct {
	Src string
	Dst string
}

type YumRepositoryParam struct {
	Config string
	Name   string
}

func (executer *Executer) Run(task Task) (HostResult, error) {

	switch task.Cmd {

	case "shell", "command":
		return executer.Shell(task)

	case "copy":
		return executer.Copy(task)

	case "template":
		return executer.Template(task)

	case "yum_repository":
		return executer.YumRepository(task)

	}

	return HostResult{

		Host:      executer.SshConfig.Host,
		Cmd:       task.Cmd,
		Result:    errors.New(fmt.Sprintf("Executer of Run Function Not Find Task.Cmd [%s]", task.Cmd)).Error(),
		StartTime: CSTLayoutString(),
		EndTime:   CSTLayoutString(),
		Status:    FAILED,
	}, errors.New("Executer of Run Function not find task.Cmd Type ")
}

/**
 * @description:
 * @param {Task} task
 * @return HostResult{},error
 */
func (executer *Executer) Shell(task Task) (HostResult, error) {

	starttime := CSTLayoutString()
	c, err := ssh.New(executer.SshConfig)
	if err != nil {

		return HostResult{

			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] When Shell Funtion  Execute, New SSH Client Error", executer.SshConfig.Host))

	}

	s := ShellParam{}

	if err = mapstructure.Decode(task.Param, &s); err != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] Shell Module Lost a Little Param", executer.SshConfig.Host))

	}

	output, err := c.Output(s.Data)

	if err != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] When Shell Funtion  Execute, Output Error", executer.SshConfig.Host))

	}

	return HostResult{
		Host:      executer.SshConfig.Host,
		Cmd:       task.Cmd,
		Result:    string(output),
		StartTime: starttime,
		EndTime:   CSTLayoutString(),
		Tag:       task.Tag,
		Status:    SUCESSFUL,
	}, nil

}

func (executer *Executer) Copy(task Task) (HostResult, error) {

	starttime := CSTLayoutString()
	c, err := ssh.New(executer.SshConfig)
	if err != nil {

		return HostResult{

			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] When Copy Funtion  Execute, New SSH Client Error", executer.SshConfig.Host))

	}

	s := CopyParam{}

	if err = mapstructure.Decode(task.Param, &s); err != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] Copy Module Lost a Little Param", executer.SshConfig.Host))

	}

	errUpload := c.Upload(s.Src, s.Dst)

	if errUpload != nil {
		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    errUpload.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errUpload

	}

	return HostResult{
		Host:      executer.SshConfig.Host,
		Cmd:       task.Cmd,
		Result:    fmt.Sprintf("Copy File[%s] Sucessful", s.Src),
		StartTime: starttime,
		EndTime:   CSTLayoutString(),
		Tag:       task.Tag,
		Status:    SUCESSFUL,
	}, nil

}

func (executer *Executer) Template(task Task) (HostResult, error) {

	starttime := CSTLayoutString()
	c, err := ssh.New(executer.SshConfig)
	if err != nil {

		return HostResult{

			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, err

	}

	s := TemplateParam{}

	if err = mapstructure.Decode(task.Param, &s); err != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, err

	}

	rendedFile, errTemplate := TemplateFile(task.FileName+"/templates/"+s.Src, task.Vars)

	if errTemplate != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errTemplate

	}

	errUpload := c.Upload(rendedFile, s.Dst)

	if errUpload != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    errUpload.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errUpload

	}

	return HostResult{
		Host:      executer.SshConfig.Host,
		Cmd:       task.Cmd,
		Result:    fmt.Sprintf("[Template] Transfer Local File: %s  Sucessful", s.Src),
		StartTime: starttime,
		EndTime:   CSTLayoutString(),
		Tag:       task.Tag,
		Status:    SUCESSFUL,
	}, nil

}

func (executer *Executer) YumRepository(task Task) (HostResult, error) {

	starttime := CSTLayoutString()
	_, err := ssh.New(executer.SshConfig)
	if err != nil {

		return HostResult{

			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] When Copy Funtion  Execute, New SSH Client Error", executer.SshConfig.Host))

	}

	s := YumRepositoryParam{}

	if err = mapstructure.Decode(task.Param, &s); err != nil {

		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, errors.Wrap(err, fmt.Sprintf("[%s] Copy Module Lost a Little Param", executer.SshConfig.Host))

	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "gosible-yum-")
	if err != nil {
		return HostResult{
			Host:      executer.SshConfig.Host,
			Cmd:       task.Cmd,
			Result:    err.Error(),
			StartTime: starttime,
			EndTime:   CSTLayoutString(),
			Tag:       task.Tag,
			Status:    FAILED,
		}, err
	}

	ioutil.WriteFile(tmpFile.Name(), []byte(s.Config), 0644)

	// errUpload := c.Upload(s.Src, s.Dst)

	// if errUpload != nil {
	// 	return HostResult{
	// 		Host:      executer.SshConfig.Host,
	// 		Cmd:       task.Cmd,
	// 		Result:    errUpload.Error(),
	// 		StartTime: starttime,
	// 		EndTime:   CSTLayoutString(),
	// 		Tag:       task.Tag,
	// 		Status:    FAILED,
	// 	}, errUpload

	// }

	return HostResult{
		Host:      executer.SshConfig.Host,
		Cmd:       task.Cmd,
		Result:    fmt.Sprintf("Copy File[%s] Sucessful", "xxxxxx"),
		StartTime: starttime,
		EndTime:   CSTLayoutString(),
		Tag:       task.Tag,
		Status:    SUCESSFUL,
	}, nil

}
