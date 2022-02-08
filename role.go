/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 16:44:06
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/gosible.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

type Role struct {
	Name  string                 `yaml:"name"`
	Vars  map[string]interface{} `yaml:"vars"`
	Tags  []string               `yaml:"tags"`
	Tasks []Task                 `yaml:tasks"`
}

type Roles []Role
