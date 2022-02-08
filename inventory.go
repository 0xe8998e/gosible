/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2022-01-01 00:22:03
 * @LastEditTime: 2022-01-28 16:44:06
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/pkg/gosible/gosible.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */
package gosible

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Host struct {
	Ip         string `yaml:"ip"`
	Port       int    `yaml:"port" default:"22"`
	UserName   string `yaml:"username"   default:"root"`
	PassWord   string `yaml:"password"`
	PrivateKey string `yaml:"private_key"`
}

// type Hosts []Host

type Group struct {
	Name  string `yaml:"name"`
	Hosts []Host `yaml:"hosts"`
}

type Groups []Group

type Inventory struct {
	Groups Groups `yaml:"groups"`
}

/**
 * @description:  Parse Invenory File Return Inventory
 * @param {string} configName
 * @return {*}
 */
func ParseInventory(configName string) Inventory {
	fileName := filepath.Base(configName)
	workDir := strings.Replace(filepath.Dir(configName), "\\", "/", -1)

	runtime_viper := viper.New()
	runtime_viper.SetConfigName(fileName)
	runtime_viper.SetConfigType("yaml")
	runtime_viper.AddConfigPath(workDir)

	if err := runtime_viper.ReadInConfig(); err != nil {

		panic(err)
	}

	inventory := Inventory{}

	if err := runtime_viper.Unmarshal(&inventory); nil != err {
		log.Fatalf("Unmarshal To Inventory Happend Error ：%v", err)
	}

	return inventory

}

/**
 * @description:
 * @param {Inventory} inventory
 * @param {string} groupName
 * @return {*}
 */
func GetHostsByGroupName(inventory Inventory, groupName string) Group {

	var group Group
	for _, v := range inventory.Groups {

		if groupName == v.Name {
			group.Name = v.Name
			group.Hosts = v.Hosts
		}
	}

	return group
}
