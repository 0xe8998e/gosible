/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 16:45:42
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/sort.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */

package gosible

type SortPlayBooks struct { //注意此处
	PlayBooks []PlayBook
	By        func(p, q *PlayBook) bool
}

func (pw SortPlayBooks) Len() int { // 重写 Len() 方法
	return len(pw.PlayBooks)
}
func (pw SortPlayBooks) Swap(i, j int) { // 重写 Swap() 方法
	pw.PlayBooks[i], pw.PlayBooks[j] = pw.PlayBooks[j], pw.PlayBooks[i]
}
func (pw SortPlayBooks) Less(i, j int) bool { // 重写 Less() 方法
	return pw.By(&pw.PlayBooks[i], &pw.PlayBooks[j])
}

func (playBooks *PlayBooks) Append(data PlayBook) *PlayBooks {
	playBooks.PlayBooks = append(playBooks.PlayBooks, data)

	return playBooks
}

// =========
type SortTasks struct { //注意此处
	Task []Task
	By   func(p, q *Task) bool
}

func (pw SortTasks) Len() int { // 重写 Len() 方法
	return len(pw.Task)
}
func (pw SortTasks) Swap(i, j int) { // 重写 Swap() 方法
	pw.Task[i], pw.Task[j] = pw.Task[j], pw.Task[i]
}
func (pw SortTasks) Less(i, j int) bool { // 重写 Less() 方法
	return pw.By(&pw.Task[i], &pw.Task[j])
}
