/*
 * @Author: 0xe8998e@gmail.com
 * @Date: 2021-09-15 14:52:26
 * @LastEditTime: 2022-02-07 22:03:18
 * @LastEditors: 0xe8998e@gmail.com
 * @FilePath: /gosible/cmd/main.go
 * @Description: gosible Contol Servers of  DevOps's  Tool
 */

package main

import (
	"fmt"
	"os"

	"github.com/0xe8998e/gosible/gosible"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "playbooks",
				Aliases:  []string{"p"},
				Value:    "./examples/playbook/test.yaml",
				Usage:    "Enter the characters",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "inventory",
				Aliases:  []string{"i"},
				Value:    "./examples/playbook/hosts.yaml",
				Usage:    "language for the greeting",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			playbooks := c.String("playbooks")
			fmt.Printf("playbooks = %v\n", playbooks)
			inventory := c.String("inventory")
			fmt.Printf("inventory = %v\n", inventory)

			gosible.InitLogger()
			Gosible := gosible.Gosible{

				Inventory: inventory,
				PlayBooks: playbooks,
			}

			Gosible.Parse()

			return nil
		},
	}

	_ = app.Run(os.Args)
}
