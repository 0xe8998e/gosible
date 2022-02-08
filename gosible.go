/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-02-07 20:34:40
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/gosible.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

import (
	"encoding/json"
)

// Define Gosible Struct As Entry
type Gosible struct {
	Inventory string `yaml:"inventory" default:"hosts"`
	PlayBooks string `yaml:"playbooks"   default:"playbooks.yml"`
}

func (gosible *Gosible) Parse() PlayBooks {

	var playBooks = PlayBooks{}

	_, perr := playBooks.ParsePlayBooks(gosible.PlayBooks)
	if perr != nil {
		Error(perr)
	}

	ierr := playBooks.ParseInventory(gosible.Inventory)
	if ierr != nil {
		Error(ierr)
	}

	r, _ := playBooks.Run()
	// 序列化
	buf, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		Error(err)
	}
	Info(string(buf))

	return playBooks

}
