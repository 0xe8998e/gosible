/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 23:01:27
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/tasks.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Tasks struct { //注意此处
	Tasks []Task `yaml:"tasks"`
}

func (tasks *Tasks) Append(data Task) *Tasks {
	tasks.Tasks = append(tasks.Tasks, data)

	return tasks
}

/**
 * @description:
 * @param {*}
 * @return {*}
 */
func (tasks *Tasks) Sort() *Tasks { // 重写 Less() 方法
	sort.Sort(SortTasks{tasks.Tasks, func(p, q *Task) bool {
		return q.Priority < p.Priority // Priority 递减排序
	}})
	return tasks
}

/**
 * @description:
 * @param {string} configName
 * @return {*}
 */
func ParseTasks(configName string, role Role) (Tasks, error) {

	tasks := Tasks{}

	playBooksRootDir := GetProjectAbsPath(configName)

	roleFile := playBooksRootDir + "/roles/" + role.Name + "/tasks/main.yaml"

	rendedFile, err := TemplateFile(roleFile, role.Vars)

	if err != nil {
		return tasks, errors.Wrap(err, fmt.Sprintf("Read Template File[%s] Error", roleFile))

	}

	fileName := filepath.Base(rendedFile)
	workDir := strings.Replace(filepath.Dir(rendedFile), "\\", "/", -1)

	// you can create a new viper instance
	// use it keep to read new config  every time
	// otherwise it will keep old data
	runtime_viper := viper.New()
	runtime_viper.SetConfigName(fileName)
	runtime_viper.SetConfigType("yaml")
	runtime_viper.AddConfigPath(workDir)

	if err := runtime_viper.ReadInConfig(); err != nil {
		// Error parsing file
		return tasks, errors.Wrap(err, fmt.Sprintf("Read In Config[%s] Happend Error", configName))
	}

	if err := runtime_viper.Unmarshal(&tasks); nil != err {

		return tasks, errors.Wrap(err, fmt.Sprintf("Unmarshal Config[%s]  Error", configName))
	}

	newTasks := Tasks{}
	for _, task := range tasks.Tasks {
		task.FileName = playBooksRootDir + "/roles/" + role.Name
		newTasks.Tasks = append(newTasks.Tasks, task)
	}
	defer os.Remove(rendedFile)

	return newTasks, nil

}
