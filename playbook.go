/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 16:51:12
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/playbook.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

type PlayBookBeforeActionFunc func(playBook *PlayBook) bool
type PlayBookAfterActionFunc func(playBook *PlayBook) bool

type PlayBook struct {
	Name          string            `yaml:"name"`
	Priority      int               `yaml:"priority"`
	GroupName     string            `yaml:"groupName"`
	Group         Group             `yaml:"group"`
	Vars          map[string]string `yaml:"vars"`
	Roles         Roles             `yaml:"roles"`
	StartDateTime string
	EndDateTime   string
	BeforeAction  []PlayBookBeforeActionFunc
	AfterAction   []PlayBookAfterActionFunc
}

func (playBook *PlayBook) SetBeforeAction(beforeAction PlayBookBeforeActionFunc) bool {

	playBook.BeforeAction = append(playBook.BeforeAction, beforeAction)
	return true
}

func (playBook *PlayBook) SetAfterAction(afterAction PlayBookAfterActionFunc) bool {

	playBook.AfterAction = append(playBook.AfterAction, afterAction)
	return true
}
