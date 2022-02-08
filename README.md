# why 

I have used ansible as the main tool of DEVOPS for a long time, but it cannot refer to the task in the role according to the specified tag in the role of the playbook, which is very uncomfortable. Later, when developing the workflow of automation tools, there is another obvious "pit" in Ansible's playbook, that is, when an error occurs in a task executed by a server, the entire playbook will not be interrupted. This is when some business lines are released. fatal flaw. I had to make some changes, so I have this gosible repository.



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