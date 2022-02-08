[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/0xe8998e/gosible?tab=doc)

# Why 

I have used ansible as the main tool of DEVOPS for a long time, but it cannot refer to the task in the role according to the specified tag in the role of the playbook, which is very uncomfortable. Later, when developing the workflow of automation tools, there is another obvious "pit" in Ansible's playbook, that is, when an error occurs in a task executed by a server, the entire playbook will not be interrupted. This is when some business lines are released. fatal flaw. I had to make some changes, so I have this gosible repository.


# Api
```
go get github.com/0xe8998e/gosible
```

```
package main

import (
	"fmt"
	"os"

	"github.com/0xe8998e/gosible"
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
```

# Build 
```
git clone github.com/0xe8998e/gosible
make buildmac|buildlinux|buildwin
```
# BINARY Run
```shell
./bin/gosible --playbooks ./examples/playbook/test.yaml  --inventory ./examples/playbook/hosts.yaml

2022-02-07T22:19:49.157+0800    INFO    gosible/gosible.go:41   {
    "playbooks": [
        {
            "name": "playbook1",
            "task_results": [
                {
                    "name": "copy plog /tmp/1.txt",
                    "host_results": [
                        {
                            "host": "10.7.180.234",
                            "cmd": "copy",
                            "result": "Copy File[/tmp/1.txt] Sucessful",
                            "start_time": "2022-02-07 22:19:44",
                            "end_time": "2022-02-07 22:19:46",
                            "tag": "test1",
                            "status": "sucessful"
                        },
                        {
                            "host": "10.7.180.234",
                            "cmd": "copy",
                            "result": "Copy File[/tmp/1.txt] Sucessful",
                            "start_time": "2022-02-07 22:19:44",
                            "end_time": "2022-02-07 22:19:46",
                            "tag": "test1",
                            "status": "sucessful"
                        },
                        {
                            "host": "10.7.180.234",
                            "cmd": "copy",
                            "result": "Copy File[/tmp/1.txt] Sucessful",
                            "start_time": "2022-02-07 22:19:44",
                            "end_time": "2022-02-07 22:19:46",
                            "tag": "test1",
                            "status": "sucessful"
                        }
                    ],
                    "start_time": "2022-02-07 22:19:44",
                    "end_time": "2022-02-07 22:19:46"
                },
                {
                    "name": "test shell 1",
                    "host_results": [
                        {
                            "host": "10.7.180.234",
                            "cmd": "shell",
                            "result": "AndroidStudioProjects\nDesktop\nDocuments\nDownloads\nLibrary\nMovies\nMusic\nPictures\nPublic\nantia\nhot.py\ntools\nwork\n",
                            "start_time": "2022-02-07 22:19:46",
                            "end_time": "2022-02-07 22:19:47",
                            "tag": "test1",
                            "status": "sucessful"
                        },
                        {
                            "host": "10.7.180.234",
                            "cmd": "shell",
                            "result": "AndroidStudioProjects\nDesktop\nDocuments\nDownloads\nLibrary\nMovies\nMusic\nPictures\nPublic\nantia\nhot.py\ntools\nwork\n",
                            "start_time": "2022-02-07 22:19:46",
                            "end_time": "2022-02-07 22:19:47",
                            "tag": "test1",
                            "status": "sucessful"
                        },
                        {
                            "host": "10.7.180.234",
                            "cmd": "shell",
                            "result": "AndroidStudioProjects\nDesktop\nDocuments\nDownloads\nLibrary\nMovies\nMusic\nPictures\nPublic\nantia\nhot.py\ntools\nwork\n",
                            "start_time": "2022-02-07 22:19:46",
                            "end_time": "2022-02-07 22:19:47",
                            "tag": "test1",
                            "status": "sucessful"
                        }
                    ],
                    "start_time": "2022-02-07 22:19:46",
                    "end_time": "2022-02-07 22:19:47"
                },
                {
                    "name": "template test",
                    "host_results": [
                        {
                            "host": "10.7.180.234",
                            "cmd": "template",
                            "result": "[Template] Transfer Local File: main.j2  Sucessful",
                            "start_time": "2022-02-07 22:19:47",
                            "end_time": "2022-02-07 22:19:49",
                            "tag": "test1",
                            "status": "sucessful"
                        },
                        {
                            "host": "10.7.180.234",
                            "cmd": "template",
                            "result": "[Template] Transfer Local File: main.j2  Sucessful",
                            "start_time": "2022-02-07 22:19:47",
                            "end_time": "2022-02-07 22:19:49",
                            "tag": "test1",
                            "status": "sucessful"
                        },
                        {
                            "host": "10.7.180.234",
                            "cmd": "template",
                            "result": "[Template] Transfer Local File: main.j2  Sucessful",
                            "start_time": "2022-02-07 22:19:47",
                            "end_time": "2022-02-07 22:19:49",
                            "tag": "test1",
                            "status": "sucessful"
                        }
                    ],
                    "start_time": "2022-02-07 22:19:47",
                    "end_time": "2022-02-07 22:19:49"
                }
            ],
            "start_time": "2022-02-07 22:19:44",
            "end_time": "2022-02-07 22:19:49"
        }
    ],
    "start_time": "2022-02-07 22:19:44",
    "end_time": "2022-02-07 22:19:49"
}
```