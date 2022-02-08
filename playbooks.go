/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-02-07 20:32:57
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/playbooks.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */

package gosible

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type PlayBooksBeforeActionFunc func(playBooks *PlayBooks) bool
type PlayBooksAfterActionFunc func(playBooks *PlayBooks) bool

type PlayBooks struct {
	Inventory    Inventory
	PlayBooks    []PlayBook `yaml:"playbooks"`
	BeforeAction []PlayBooksBeforeActionFunc
	AfterAction  []PlayBooksAfterActionFunc
}

func (playBooks *PlayBooks) Sort() *PlayBooks { // 重写 Less() 方法
	sort.Sort(SortPlayBooks{playBooks.PlayBooks, func(p, q *PlayBook) bool {
		return q.Priority < p.Priority // Priority 递减排序
	}})
	return playBooks
}

func (playBooks *PlayBooks) SetBeforeAction(beforeAction PlayBooksBeforeActionFunc) bool {

	playBooks.BeforeAction = append(playBooks.BeforeAction, beforeAction)
	return true
}

func (playBooks *PlayBooks) SetAfterAction(afterAction PlayBooksAfterActionFunc) bool {

	playBooks.AfterAction = append(playBooks.AfterAction, afterAction)
	return true
}

func (playBooks *PlayBooks) ParseInventory(configName string) error {
	fileName := filepath.Base(configName)
	workDir := strings.Replace(filepath.Dir(configName), "\\", "/", -1)

	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	inventory := Inventory{}

	if err := viper.Unmarshal(&inventory); nil != err {
		Errorf("Unmarshal To Inventory  Error ：%v", err)
		return err

	}

	playBooks.Inventory = inventory

	return nil

}

func (playBooks *PlayBooks) ParsePlayBooks(configName string) (*PlayBooks, error) {
	fileName := filepath.Base(configName)
	workDir := strings.Replace(filepath.Dir(configName), "\\", "/", -1)

	runtime_viper := viper.New()
	runtime_viper.SetConfigName(fileName)
	runtime_viper.SetConfigType("yaml")
	runtime_viper.AddConfigPath(workDir)

	if err := runtime_viper.ReadInConfig(); err != nil {

		return playBooks, errors.Wrap(err, fmt.Sprintf("Read Config File[%s] Failed", configName))

	}

	// playBooksData := PlayBooks{}

	if err := runtime_viper.Unmarshal(&playBooks); nil != err {
		return playBooks, errors.Wrap(err, fmt.Sprintf("Unmarshal Config File[%s] Failed", configName))

	}

	for _, playBook := range playBooks.PlayBooks {

		temp := make([]Role, 0)
		for _, role := range playBook.Roles {

			tasks, err := ParseTasks(configName, role)

			if err != nil {
				Error(err)
				break
			}

			tasksFilterTagsResult := make([]Task, 0)
			for _, task := range tasks.Tasks {
				if InStringSlice(role.Tags, task.Tag) {
					tasksFilterTagsResult = append(tasksFilterTagsResult, task)
				}
			}

			role.Tasks = tasksFilterTagsResult

			temp = append(temp, role)
		}

		copier.Copy(&playBook.Roles, &temp)

		// Infof("%#v", playBook.Roles)

	}

	return playBooks, nil

}

func (playBooks *PlayBooks) Run() (PlayBooksResult, error) {

	var (
		runFlag = false
	)

	for _, cb := range playBooks.BeforeAction {

		cb(playBooks)
	}

	playBooksResult := PlayBooksResult{}
	playBooksResult.StartTime = CSTLayoutString()

	for _, playBook := range playBooks.PlayBooks {

		playBookResult := PlayBookResult{}
		playBookResult.Name = playBook.Name
		playBookResult.StartTime = CSTLayoutString()

		group := GetHostsByGroupName(playBooks.Inventory, playBook.GroupName)

		for _, role := range playBook.Roles {

			for _, task := range role.Tasks {

				task.Vars = role.Vars
				taskResult := TaskResult{}
				taskResult.Name = task.Name
				taskResult.StartTime = CSTLayoutString()

				done := make(chan interface{})
				defer close(done)
				inStream := make(chan interface{})

				go func() {
					defer close(inStream)
					for i := 0; i < len(group.Hosts); i++ {

						xx := ExecuterParam{
							Host: group.Hosts[i],
							Task: task,
						}

						inStream <- xx
					}
				}()

				// 扇入扇出执行耗时任务
				resultChan := fanIn(done, fanOut(done, inStream)...)

				hostResults := make([]HostResult, 0)
				// 使用chRange安全遍历打印
				for val := range resultChan {

					hostResults = append(hostResults, val.(ExecuterResult).HostResult)

					if val.(ExecuterResult).Error != nil {

						runFlag = true
					}

				}

				// hostResults := make([]HostResult, 0)
				// for _, host := range group.Hosts {

				// 	executer := Executer{
				// 		SshConfig: &ssh.Config{
				// 			Host: host.Ip,
				// 			Port: host.Port,
				// 			User: host.UserName,
				// 		},
				// 	}

				// 	task.Vars = role.Vars
				// 	c, err := executer.Run(task)
				// 	hostResults = append(hostResults, c)
				// 	if err != nil {

				// 		continue
				// 	}

				// }

				taskResult.HostResults = hostResults
				taskResult.EndTime = CSTLayoutString()

				playBookResult.TaskResults = append(playBookResult.TaskResults, taskResult)

				if runFlag {

					break
				}
			}
			if runFlag {
				break
			}

		}

		playBookResult.EndTime = CSTLayoutString()
		playBooksResult.PlayBookResults = append(playBooksResult.PlayBookResults, playBookResult)
		if runFlag {
			break
		}
	}
	playBooksResult.EndTime = CSTLayoutString()

	for _, cb := range playBooks.AfterAction {

		cb(playBooks)
	}

	return playBooksResult, nil
}
