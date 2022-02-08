/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 22:17:00
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/task.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

type TaskBeforeActionFunc func(task *Task) bool
type TaskAfterActionFunc func(task *Task) bool

type Task struct {
	Id            string
	Name          string
	Cmd           string
	FileName      string
	Tag           string
	Vars          Vars
	Param         Param
	Priority      int
	StartDateTime string
	EndDateTime   string
	BeforeAction  []TaskBeforeActionFunc
	AfterAction   []TaskAfterActionFunc
}

func (task *Task) SetBeforeAction(beforeAction TaskBeforeActionFunc) bool {

	task.BeforeAction = append(task.BeforeAction, beforeAction)
	return true
}

func (task *Task) SetAfterAction(afterAction TaskAfterActionFunc) bool {

	task.AfterAction = append(task.AfterAction, afterAction)
	return true
}
