/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 20:43:30
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/result.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

type PlayBooksResult struct {
	PlayBookResults []PlayBookResult `json:"playbooks"`
	StartTime       string           `json:"start_time"`
	EndTime         string           `json:"end_time"`
}

type PlayBookResult struct {
	Name        string       `json:"name"`
	TaskResults []TaskResult `json:"task_results"`
	StartTime   string       `json:"start_time"`
	EndTime     string       `json:"end_time"`
}

type TaskResult struct {
	Name        string       `json:"name"`
	HostResults []HostResult `json:"host_results"`
	StartTime   string       `json:"start_time"`
	EndTime     string       `json:"end_time"`
}

type HostResult struct {
	Host      string `json:"host"`
	Cmd       string `json:"cmd"`
	Result    string `json:"result"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Tag       string `json:"tag"`
	Status    string `json:"status"`
}
